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

// API URL
// For more documentation on the API visit: https://twitchappapi.docs.apiary.io/
const API = "https://addons-ecs.forgesvc.net/api/v2/"

// Globals
var (
	jarFingerprints map[int]string      // Contains 32-bit MurmurHash2 of each .jar and the fileName
	oldMap          map[string][]string // Old mods map
	newMap          map[string][]string // New mods map
	externalMods    []string            // Mods that cannot be found on CurseForge
	currentTime     time.Time           // Current time
)

// Command line Arguments
var (
	instancePath       *string // Path to instance folder
	gameVersion        *string // Specified game version
	downloadPath       *string // Path for new .jar files
	exportNewManifest  *bool   // Export manifest.json
	exportOldManifest  *bool   // Export oldmanifest.json
	exportManifestPath *string // Path to export.json
	silentMode         *bool   // Silent mode
	cliFlag            *bool   // Enable CLI
	guiFlag            *bool   // Enable GUI
)

func main() {

	parseArgs()
	getTime()

	if *cliFlag {
		parseArgs()
		readInstancePath()
		useArgs()
		checkUpdates(oldMap, newMap)
	}

	if *guiFlag {
		launchGUI()
		readInstancePath()
		useArgs()
		checkUpdates(oldMap, newMap)
	}
}

// Parse arguments specified in the CLI
func parseArgs() {

	cliFlag = flag.Bool("cli", false, "CLI. Must be followed up by arguments.")
	guiFlag = flag.Bool("gui", false, "GUI. Standalone, don't bother using additional arguments.")

	instancePath = flag.String("d", "./", "Absolute path to Minecraft instance folder.")
	gameVersion = flag.String("version", "1.12.2", "Game version of located mods.")
	downloadPath = flag.String("download", "./", "Path of where to download .jar file")

	exportNewManifest = flag.Bool("export-new", false, "Creation of new manifest.json")
	exportOldManifest = flag.Bool("export-old", false, "Creation of old manifest.json")
	exportManifestPath = flag.String("manifest", "./", "Absolute path of export.json")

	silentMode = flag.Bool("s", false, "Silent mode.")

	flag.Parse()

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

// Determine if instance folder is valid & reads it
func readInstancePath() {

	err := readInstanceFolder(*instancePath)

	if err != nil {
		log.Fatal(err)
	}

}

// Use arguments specified in the CLI
func useArgs() {

	if *exportNewManifest {
		readExport(*exportManifestPath, "new")
	}

	if *exportOldManifest {
		readExport(*exportManifestPath, "old")
	}

	if *silentMode {
		log.SetOutput(ioutil.Discard)
	}

}

// Actually read the instance folder
func readInstanceFolder(path string) error {

	log.Println("Reading Minecraft directory...")

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Println("Cannot read directory.")
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.ToLower(f.Name()) == "mods" {
			log.Println("Reading Minecraft directory completed.")
			listMods(path)
			return nil
		}
	}

	return errors.New("Cannot find mods folder.")

}

func listMods(modsFolder string) {

	log.Print("Reading mods folder...")

	jarFingerprints = make(map[int]string)

	err := filepath.Walk(path.Join(modsFolder, "mods"), func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".jar" {
			fileHash, _ := GetFileHash(path)
			jarFingerprints[fileHash] = info.Name()
		}
		return nil
	})

	if err != nil {
		log.Println("Error searching through mods folder.")
		log.Fatal(err)
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

	log.Printf("Mods that can't be found on CurseForge: %v", externalMods)

}
