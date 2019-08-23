package main

// Manifest contains all contents within the .json file
type Manifest struct {
	Minecraft       Minecraft  `json:"minecraft"`
	ManifestType    string     `json:"manifestType"`
	ManifestVersion int        `json:"manifestVersion"`
	Name            string     `json:"name"`
	Version         string     `json:"version"`
	Author          string     `json:"author"`
	Files           []CurrFile `json:"files"`
}

// CurrFile contains all contents within each .jar file
type CurrFile struct {
	ProjectID int  `json:"projectID"`
	FileID    int  `json:"fileID"`
	Required  bool `json:"required"`
}

// Minecraft contains the version and the modloaders for the manifest
type Minecraft struct {
	Version    string       `json:"version"`
	ModLoaders []ModLoaders `json:"modLoaders"`
}

// ModLoaders contains the id and type for Minecraft
type ModLoaders struct {
	ID      string `json:"id"`
	Primary bool   `json:"primary"`
}

// ExportJSON contains all contents for the export.json file
type ExportJSON struct {
	MinecraftVersion string `json:"MinecraftVersion"`
	Modloader        string `json:"Modloader"`
	ModloaderVersion string `json:"ModloaderVersion"`
	ManifestType     string `json:"ManifestType"`
	ManifestVersion  int    `json:"ManifestVersion"`
	PackName         string `json:"PackName"`
	PackVersion      string `json:"PackVersion"`
	PackAuthors      string `json:"PackAuthors"`
}
