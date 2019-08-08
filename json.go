package main

import (
	"log"
	"math"
	"time"

	"github.com/buger/jsonparser"
)

func parseOldJSON(body []byte) (string, []string) {

	paths := [][]string{
		[]string{"id"},                  // projectID
		[]string{"file", "fileName"},    // fileName
		[]string{"file", "downloadUrl"}, // downloadURL
		[]string{"file", "id"},          // fileID
	}

	var projectID, fileName, downloadURL, fileID string

	// Traverse each element of exactMatches array
	_, err := jsonparser.ArrayEach(body, func(file []byte, dataType jsonparser.ValueType, offset int, err error) {

		// Traverse each key of file object given paths
		jsonparser.EachKey(file, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				projectID = string(value)
			case 1:
				fileName = string(value)
			case 2:
				downloadURL = string(value)
			case 3:
				fileID = string(value)
			}
		}, paths...)

	}, "exactMatches")

	if err != nil {
		log.Println("No exactMatches array found for given JSON.")
	}

	return projectID, []string{fileName, downloadURL, fileID}

}

func parseNewJSON(file []byte) []string {

	paths := [][]string{
		[]string{"fileName"},    // fileName
		[]string{"downloadUrl"}, // downloadURL
		[]string{"id"},          // fileID
	}

	var fileName, downloadURL, fileID string

	// Traverse the one object passed in
	jsonparser.EachKey(file, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		switch idx {
		case 0:
			fileName = string(value)
		case 1:
			downloadURL = string(value)
		case 2:
			fileID = string(value)
		}
	}, paths...)

	return []string{fileName, downloadURL, fileID}

}

// currentVersionExists returns the existence of (USER_VERSION) within the
// (versions) array.
// Useful for findBestFile in filtering out unnecessary objects.
func currentVersionExists(versions []string) bool {
	for i := range versions {
		if versions[i] == *USER_VERSION {
			return true
		}
	}
	return false
}

func findBestFile(body []byte) []byte {

	lowestDifference := math.MaxFloat64 // Lowest difference so far
	var searchedFile []byte             // File object found so far

	// Traverse array
	_, err := jsonparser.ArrayEach(body, func(file []byte, dataType jsonparser.ValueType, offset int, err error) {

		var versions []string // Versions found within object

		// Traverse each object within the array
		_, errVersion := jsonparser.ArrayEach(file, func(gameVersion []byte, dataType jsonparser.ValueType, offset int, err error) {
			versions = append(versions, string(gameVersion))
		}, "gameVersion")

		if errVersion != nil {
			log.Println("gameVersion array does not exist.")
		}

		// Determine if the object needs to be checked, if it does then it is parsed
		if currentVersionExists(versions) {

			fileDate_str, _ := jsonparser.GetString(file, "fileDate")
			fileDate_RFC3339Nano, _ := time.Parse(time.RFC3339Nano, fileDate_str)
			difference := currentTime.Sub(fileDate_RFC3339Nano).Minutes()

			if difference < lowestDifference {
				lowestDifference = difference
				searchedFile = file
			}

		}

	})

	if err != nil {
		log.Println("Blank array.")
	}

	return searchedFile

}
