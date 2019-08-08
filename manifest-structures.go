package main

type Manifest struct {
	Minecraft       Minecraft  `json:"minecraft"`
	ManifestType    string     `json:"manifestType"`
	ManifestVersion int        `json:"manifestVersion"`
	Name            string     `json:"name"`
	Version         string     `json:"version"`
	Author          string     `json:"author"`
	Files           []CurrFile `json:"files"`
}

type CurrFile struct {
	ProjectID string `json:"projectID"`
	FileID    string `json:"fileID"`
	Required  bool   `json:"required"`
}

type Minecraft struct {
	Version    string       `json:"version"`
	ModLoaders []ModLoaders `json:"modLoaders"`
}

type ModLoaders struct {
	Id      string `json:"id"`
	Primary bool   `json:"primary"`
}

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
