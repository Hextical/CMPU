package main

import (
	"log"
	"sync"
)

func createMaps(jarFingerprints map[int]string) {
	createOldMap(jarFingerprints)
	createNewMap(oldMap)
}

func createOldMap(jarFingerprints map[int]string) {

	log.Println("Creating old map...")

	oldMap = make(map[string][]string)

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for key, value := range jarFingerprints {
		wg.Add(1)
		go func(wg *sync.WaitGroup, key int, value string) {
			defer wg.Done()
			tempBody := connectWithHash(key)
			projectID, attributes := parseOldJSON(tempBody)

			if projectID == "" { // Can't find a projectID -> Add to an external array
				mutex.Lock()
				externalMods = append(externalMods, value)
				mutex.Unlock()
			} else { // Otherwise set the key & value to the map
				mutex.Lock()
				oldMap[projectID] = attributes
				mutex.Unlock()
			}

		}(&wg, key, value)
	}

	wg.Wait()

	log.Println("Old map created.")

}

func createNewMap(oldMap map[string][]string) {

	log.Println("Creating new map...")

	newMap = make(map[string][]string)

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for key := range oldMap {
		wg.Add(1)
		go func(wg *sync.WaitGroup, key string) {
			defer wg.Done()
			tempBody := connectWithProjectID(key)
			attributes := parseNewJSON(findBestFile(tempBody))
			mutex.Lock()
			newMap[key] = attributes
			mutex.Unlock()
		}(&wg, key)
	}

	wg.Wait()

	log.Println("New map created.")

}
