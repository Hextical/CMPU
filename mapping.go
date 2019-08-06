package main

import (
	"log"
	"sync"
)

func createMaps(jarFingerprints []int) {
	createOldMap(jarFingerprints)
	createNewMap(oldMap)
}

func createOldMap(jarFingerprints []int) {

	log.Println("Creating old map...")

	oldMap = make(map[string][]string)

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := range jarFingerprints {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			tempBody := connectWithHash(jarFingerprints[i])
			projID, attributes := parseOldJSON(tempBody)
			mutex.Lock()
			oldMap[projID] = attributes
			mutex.Unlock()
		}(&wg, i)
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
