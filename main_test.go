package main

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {

	getTime()

	var mcDir *string
	mcDirStr := "C:\\Users\\hexii\\Documents\\MultiMC\\instances\\1.12.2\\.minecraft"
	mcDir = &mcDirStr

	USER_VERSION_STR := "1.12.2"
	USER_VERSION = &USER_VERSION_STR

	err := readMCDIR(*mcDir)

	if err != nil {
		log.Fatal("Cannot find mods folder.")
	}

	downloadPath_str := "C:\\Users\\hexii\\Downloads\\CMPU-Downloads"
	downloadPath = &downloadPath_str

	checkUpdates(oldMap, newMap)

}

func TestExport(t *testing.T) {

	USER_VERSION_STR := "1.12.2"
	USER_VERSION = &USER_VERSION_STR

	EXPORT_NEW_BOOL := true
	EXPORT_NEW := &EXPORT_NEW_BOOL
	EXPORT_PATH_STR := "C:\\Users\\hexii\\go\\src\\CMPU\\export.json"
	EXPORT_PATH := &EXPORT_PATH_STR

	if *EXPORT_NEW {
		readExport(*EXPORT_PATH, "new")

	}

}
