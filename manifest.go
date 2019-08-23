package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/buger/jsonparser"
)

func readExport(exportPath string, manifestType string) {

	log.Println("Reading export.json...")

	export, err := ioutil.ReadFile(exportPath)

	if err != nil {
		log.Println(err)
		log.Println("***WARNING: export.json is missing, either fill in the json manually",
			"or retry with a correct export.json***")
	}

	json := readExportJSON(export)

	if manifestType == "old" {
		manifest(json, oldMap, manifestType)
	} else {
		manifest(json, newMap, manifestType)
	}

	log.Println("Reading export.json completed.")

}

func readExportJSON(file []byte) ExportJSON {

	paths := [][]string{
		{"MinecraftVersion"},
		{"Modloader"},
		{"ModloaderVersion"},
		{"ManifestType"},
		{"ManifestVersion"},
		{"PackName"},
		{"PackVersion"},
		{"PackAuthors"},
	}

	var json ExportJSON

	jsonparser.EachKey(file, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		switch idx {
		case 0:
			json.MinecraftVersion = string(value)
		case 1:
			json.Modloader = string(value)
		case 2:
			json.ModloaderVersion = string(value)
		case 3:
			json.ManifestType = string(value)
		case 4:
			json.ManifestVersion, _ = strconv.Atoi(string(value))
		case 5:
			json.PackName = string(value)
		case 6:
			json.PackVersion = string(value)
		case 7:
			json.PackAuthors = string(value)
		}
	}, paths...)

	return json

}

func manifest(exportjson ExportJSON, xMap map[string][]string, manifestType string) {

	var moddedFiles []CurrFile

	for key, value := range xMap {
		keyInt, _ := strconv.Atoi(key)
		valueInt, _ := strconv.Atoi(value[2])
		moddedFiles = append(moddedFiles, CurrFile{keyInt, valueInt, true})
	}

	data := Manifest{
		Minecraft: Minecraft{
			Version: *gameVersion,
			ModLoaders: []ModLoaders{
				{
					ID:      exportjson.Modloader + "-" + exportjson.ModloaderVersion,
					Primary: true,
				},
			},
		},
		ManifestType:    exportjson.ManifestType,
		ManifestVersion: exportjson.ManifestVersion,
		Name:            exportjson.PackName,
		Version:         exportjson.PackVersion,
		Author:          exportjson.PackAuthors,
		Files:           moddedFiles,
	}

	file, _ := json.MarshalIndent(data, "", "  ")

	if manifestType == "old" {
		_ = ioutil.WriteFile("old.json", file, 0644)
	} else {
		_ = ioutil.WriteFile("manifest.json", file, 0644)
	}

}
