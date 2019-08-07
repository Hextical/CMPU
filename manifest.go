package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/buger/jsonparser"
)

func readExport(EXPORT_PATH string) {

	log.Println("Reading export json...")

	export, err := ioutil.ReadFile(EXPORT_PATH)

	if err != nil {
		log.Println(err)
	}

	json := readExportJSON(export)
	manifest(json)

	log.Println("Reading export json completed.")

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

	// yeah this is kinda bad since I'm returning the whole structure
	return json

}

func manifest(exportjson ExportJSON) {
	data := File{
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
	}

	file, _ := json.MarshalIndent(data, "", "  ")

	_ = ioutil.WriteFile("test.json", file, 0644)

}
