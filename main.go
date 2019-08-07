package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Globals
var (
	jarFingerprints []int               // Contains 32-bit MurmurHash2 of each .jar
	oldMap          map[string][]string // Old mods map
	newMap          map[string][]string // New mods map
	USER_VERSION    *string             // Specified game version
	currentTime     time.Time           // Current time
)

func main() {
	getTime()
	runArgs()
}

// getTime gets the current time and stores it into currentTime as a
// RFC3339Nano format.
// Example: currentTime = 2019-08-06 16:22:02.5950613 -0400 EDT
func getTime() {

	log.Println("Retrieving current local time...")

	var err error

	currentTime_str := time.Now().Format(time.RFC3339Nano)
	currentTime, err = time.Parse(time.RFC3339Nano, currentTime_str)

	if err != nil {
		log.Println("Cannot parse current time/date.")
		log.Fatal(err)
	}

	log.Printf("Current local time: %v", currentTime)

}

func runArgs() {

	mcDir := flag.String("d", "./", "Absolute path to Minecraft instance folder.")
	USER_VERSION = flag.String("version", "1.12.2", "Game version of located mods.")

	EXPORT_NEW := flag.String("export", "false", "Creation of new manifest.json")
	EXPORT_PATH := flag.String("export-path", "./", "Path of manifest.json")

	flag.Parse()

	err := readMCDIR(*mcDir)

	if err != nil {
		log.Fatal(err)
	}

	if *EXPORT_NEW == "true" {
		readExport(*EXPORT_PATH)
	}

	checkUpdates(oldMap, newMap)

}

func readMCDIR(dirPath string) error {

	log.Println("Reading Minecraft directory...")

	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		log.Println("Cannot read directory.")
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.ToLower(f.Name()) == "mods" {
			log.Println("Reading Minecraft directory completed.")
			listMods(dirPath)
			return nil
		}
	}

	return errors.New("Cannot find mods folder.")

}

func listMods(modsFolder string) {

	log.Print("Reading mods folder...")

	files, err := ioutil.ReadDir(path.Join(modsFolder, "mods"))

	if err != nil {
		log.Println("Cannot read mods folder.")
		log.Fatal(err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".jar" {
			fileHash, _ := GetFileHash(path.Join(modsFolder, "mods", f.Name()))
			jarFingerprints = append(jarFingerprints, fileHash)
		}
	}

	log.Println("Reading mods folder completed.")

	createMaps(jarFingerprints)

}

func checkUpdates(oldMap map[string][]string, newMap map[string][]string) {

	for key, value := range oldMap {
		fileID_OLD := value[2]
		fileID_NEW := newMap[key][0]
		if fileID_OLD != fileID_NEW {
			log.Println(value)
		}
	}

}
