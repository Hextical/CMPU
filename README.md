# CMPU (_Curse Modpack Utilities_)

[![Go Report Card](https://goreportcard.com/badge/github.com/Hextical/CMPU)](https://goreportcard.com/report/github.com/Hextical/CMPU)

**Currently a work in progress**

Please report any issues to the issues tab with a log of what commands were used and the output. 

## Usage
Full CLI usage so far:

`
CMPU -cli=true -d C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft -version 1.12.2 -export-new=true -export-old=true -manifest C:\Users\hexii\Desktop\export.json -download C:\Users\hexii\Desktop\CMPU-Downloads
`

**Recommended: USE QUOTES FOR PATHS, ESPECIALLY IF IT CONTAINS SPACES**

Example: 

`CMPU -cli=true -d C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft`

turns into

`CMPU -cli=true -d "C:\Users\hexii\Documents\MultiMC\instances\1.12.2\.minecraft"`

Explanation:

Argument | What it does | Default value | Options | Extra info

If no value for an argument is specified it will use the default.

- `-gui=<boolean>` | graphics user interface, if specified do not use any arguments below | `false`
- `-cli=<boolean>` | command line interface, must be followed up by some of the arguments below | `false`
- `-d <path>` | instance folder path | `./`
- `-version <gameversion>` | game version | `1.12.2`
- `-release <string>` | release type | `stable` | options: stable, latest | must be lowercase & any misspelling = latest
- `-export-new=<boolean>` | if an updated manifest should be generated | `false`
- `-export-old=<boolean>` | if an old manifest should be generated | `false`
- `-manifest <path>` | path to export.json | `./`
- `-download <path>` | path for downloading updates; **it will create a folder if it does not exist** | `./`
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
Navigate to directory and run `go build`. Requires [jsonparser](https://github.com/buger/jsonparser), [gotk3](https://github.com/gotk3/gotk3), and [go-murmur](https://github.com/aviddiviner/go-murmur).

## Planned features
- View the [projects tab](https://github.com/Hextical/CMPU/projects).
