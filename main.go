package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
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
	downloadPath    *string             // Path for new .jar files
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

	EXPORT_NEW := flag.Bool("export-new", false, "Creation of new manifest.json")
	EXPORT_OLD := flag.Bool("export-old", false, "Creation of old manifest.json")
	EXPORT_PATH := flag.String("manifest", "./", "Absolute path of manifest.json")

	downloadPath = flag.String("download", "./", "Path of where to download .jar file")

	flag.Parse()

	err := readMCDIR(*mcDir)

	if err != nil {
		log.Fatal(err)
	}

	if *EXPORT_OLD {
		readExport(*EXPORT_PATH, "old")
	}

	if *EXPORT_NEW {
		readExport(*EXPORT_PATH, "new")
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

	// Create the directory if downloads requested
	if *downloadPath != "./" {

		err := os.Mkdir(*downloadPath, os.ModePerm)

		if err != nil {
			log.Println(err)
		}

	}

	log.Println("--- Updates ---")

	var updates int

	for key, value := range newMap {

		if value[2] != oldMap[key][2] { // (new fileID) != (old fileID)

			log.Printf("Name: %v | URL: %v | ID: %v", value[0], value[1], value[2])
			updates++

			// Actually download the files
			if *downloadPath != "./" {

				// Pass in the fileName and downloadUrl
				err := DownloadFile(value[0], value[1])

				if err != nil {
					log.Println(err)
				}

			}

		}
	}

	log.Println("--- End of updates ---")
	log.Printf("* Available updates: %v *", updates)

}
