# CMPU (_Curse Modpack Utilities_)

**EXTREMELY BETA**

Please report any issues to the issues tab with a log of what commands were used and the output. 

This is my first time ever using Go and I'm not a CS student, so suggestions and enhancements are greatly appreciated.

Note: This project is a rewrite of my old [updater](https://github.com/Hextical/updater-java/).
Currently this can be used to check for updates from a given `instances` folder. Only works with CurseForge mods right now.

## Usage

One can utilize the program by locating the absolute path to the `instance` folder. Check out the releases tab for a binary.

Example command-line usage:

`
CMPU -d "C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft" "1.12.2"
`

Full usage so far (checks for updates, then exports both the old and new manifest json files)

`
CMPU -d "C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft" "1.12.2" -export-new true -export-old true -manifest C:\Users\hexii\Desktop\export.json
`

This will generate two files: an old.json and a manifest.json file within the execution directory.

To see all possible commands use:

`
CMPU -help
`

## Building
Navigate to directory and run `go build`.

## Requirements
- [fasthttp](https://github.com/valyala/fasthttp)
- [jsonparser](https://github.com/buger/jsonparser)

## Planned features
- Exporting a manifest file so one can utilize the [ChangelogGenerator](https://github.com/TheRandomLabs/ChangelogGenerator)
- Creating my own ChangelogGenerator
- Exporting a complete modpack (yes it's been done before)
- Downloading the updates
