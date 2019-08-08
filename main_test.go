package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	// This just sets the annoying stuff to default
	parseArgs()

	// Actually set test the updater & downloader
	instancePathStr := "C:\\Users\\hexii\\Documents\\MultiMC\\instances\\1.12.2\\.minecraft"
	instancePath = &instancePathStr

	gameVersionStr := "1.12.2"
	gameVersion = &gameVersionStr

	downloadPath_str := "C:\\Users\\hexii\\Downloads\\CMPU-Downloads"
	downloadPath = &downloadPath_str

	getTime()
	readInstancePath()
	useArgs()
	checkUpdates(oldMap, newMap)
}

func TestExportNewManifest(t *testing.T) {

	gameVersionStr := "1.12.2"
	gameVersion = &gameVersionStr

	exportNewManifestBool := true
	exportNewManifest := &exportNewManifestBool

	exportManifestPathStr := "C:\\Users\\hexii\\go\\src\\CMPU\\export.json"
	exportManifestPath := &exportManifestPathStr

	if *exportNewManifest {
		readExport(*exportManifestPath, "new")
	}

}
