# CMPU (_Curse Modpack Utilities_)

**EXTREMELY BETA**

Please report any issues to the issues tab with a log of what commands were used and the output. 

This is my first time ever using Go and I'm not a CS student, so suggestions and enhancements are greatly appreciated.

Note: This project is a rewrite of my old [updater](https://github.com/Hextical/updater-java/).
Currently this can be used to check for updates from a given `instances` folder. Only works with CurseForge mods right now.

## Usage
Full CLI usage so far:

`
CMPU -d C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft -version 1.12.2 -export-new=true -export-old=true -manifest C:\Users\hexii\Desktop\export.json -download C:\Users\hexii\Desktop\CMPU-Downloads
`

**Recommended: USE QUOTES FOR PATHS, ESPECIALLY IF IT CONTAINS SPACES**

Example: `CMPU -d C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft` turns into `CMPU -d "C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft"`

Explanation:

Argument | What it does | Default value

If no value for an argument is specified it will use the default.

- `-d <path> ` | instance folder | `./`
- `-version <gameversion> ` | game version | `1.12.2`
- `-export-new=<boolean>` | if an updated manifest should be generated | `false`
- `-export-old=<boolean>` | if an old manifest should be generated | `false`
- `-manifest <path>` | path to export.json | `./`
- `-download <path>` | path for downloading updates; it will create a folder if it does not exist | `./`
- `-s=<boolean>` | silent mode; no output for CLI | `false`

`export.json`:

``` json
{
    "MinecraftVersion": "1.12.2",
    "Modloader": "forge",
    "ModloaderVersion": "14.23.5.2838",
    "ManifestType": "minecraftModpack",
    "ManifestVersion": 1,
    "PackName": "Example Pack",
    "PackVersion": "1.0.0",
    "PackAuthors": "author1, author2"
}
```

To see all possible commands use:

`
CMPU -help
`

## Building
Navigate to directory and run `go build`. Requires [jsonparser](https://github.com/buger/jsonparser) to be installed.

## Planned features
- View the [projects tab](https://github.com/Hextical/CMPU/projects).
