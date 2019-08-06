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

	checkUpdates(oldMap, newMap)

}
