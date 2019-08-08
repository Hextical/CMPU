package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/buger/jsonparser"
)

func readExport(EXPORT_PATH string, manifestType string) {

	log.Println("Reading export.json...")

	export, err := ioutil.ReadFile(EXPORT_PATH)

	if err != nil {
		log.Println(err)
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
		[]string{"MinecraftVersion"},
		[]string{"Modloader"},
		[]string{"ModloaderVersion"},
		[]string{"ManifestType"},
		[]string{"ManifestVersion"},
		[]string{"PackName"},
		[]string{"PackVersion"},
		[]string{"PackAuthors"},
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

func manifest(exportjson ExportJSON, x_map map[string][]string, manifestType string) {

	var moddedFiles []CurrFile

	for key, value := range x_map {
		moddedFiles = append(moddedFiles, CurrFile{key, value[2], true})
	}

	data := Manifest{
		Minecraft: Minecraft{
			Version: *USER_VERSION,
			ModLoaders: []ModLoaders{
				ModLoaders{
					Id:      exportjson.Modloader + "-" + exportjson.ModloaderVersion,
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
