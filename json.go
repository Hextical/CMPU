package main

import (
	"log"
	"math"
	"time"

	"github.com/buger/jsonparser"
)

func parseOldJSON(body []byte) (string, []string) {

	paths := [][]string{
		{"id"},                  // projectID
		{"file", "fileName"},    // fileName
		{"file", "downloadUrl"}, // downloadURL
		{"file", "id"},          // fileID
	}

	// Check if the mod was found on CurseForge
	_, _, _, emptyerr := jsonparser.Get(body, "exactMatches", "[0]")
	if emptyerr != nil {
		return "", []string{"", "", ""}
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
		{"fileName"},    // fileName
		{"downloadUrl"}, // downloadURL
		{"id"},          // fileID
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

// filterGameVersion filters the versions based on the user's specification.
// Useful for findBestFile in filtering out unnecessary objects.
func filterGameVersion(versions []string) bool {
	for i := range versions {
		if versions[i] == *gameVersion {
			return true
		}
	}
	return false
}

// filterReleaseType filters the file type based on the user's specification
// and returns if the specification exists within that file.
func filterReleaseType(file []byte, releaseIteration int) bool {

	// Don't check anything if it's latest
	if *releaseType == "latest" {
		return true
	}

	// Grab releaseType from file
	releaseKey, err := jsonparser.GetInt(file, "releaseType")

	if err != nil {
		log.Println("Unable to find releaseType", err)
	}

	if *releaseType == "stable" {
		// 1 = release, 2 = beta, 3 = alpha
		if int(releaseKey) == releaseIteration {
			return true
		}
		return false
	}

	// Not supposed to get here, but let's just return the latest
	return true

}

func findBestFile(body []byte, releaseIteration int) []byte {

	lowestDifference := math.MaxFloat64 // Lowest difference so far
	var searchedFile []byte             // File object found so far

	// Traverse array of files
	_, err := jsonparser.ArrayEach(body, func(file []byte, dataType jsonparser.ValueType, offset int, err error) {

		var versions []string // Versions found within object

		// Traverse each object within the array
		_, errVersion := jsonparser.ArrayEach(file, func(gameVersion []byte, dataType jsonparser.ValueType, offset int, err error) {
			versions = append(versions, string(gameVersion))
		}, "gameVersion")

		if errVersion != nil {
			log.Println("gameVersion array does not exist.")
		}

		// Determine if the object needs to be checked
		if filterGameVersion(versions) {
			// Store an old temp variable of the current file

			if filterReleaseType(file, releaseIteration) {

				fileDateStr, _ := jsonparser.GetString(file, "fileDate")
				fileDateRFC3339Nano, _ := time.Parse(time.RFC3339Nano, fileDateStr)
				difference := currentTime.Sub(fileDateRFC3339Nano).Minutes()

				if difference < lowestDifference {
					lowestDifference = difference
					searchedFile = file
				}

			}

		}

	})

	if err != nil {
		log.Println("Blank array.")
	}

	return searchedFile

}
