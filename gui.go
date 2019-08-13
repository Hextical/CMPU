package main

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

// Setter for all the values in the GUI
func setFromGUI(ipStr string, gvStr string, dpStr string, empStr string) {
	instancePath = &ipStr
	gameVersion = &gvStr
	downloadPath = &dpStr
	exportManifestPath = &empStr
}

// Annoyingly huge function, could be broken up
func makeBasicControlsPage() ui.Control {

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	// All the entries contents
	instancePathContents := ui.NewEntry()
	gameVersionContents := ui.NewEntry()
	downloadPathContents := ui.NewEntry()
	exportManifestPath := ui.NewEntry()

	// Default values
	instancePathContents.SetText("./")
	gameVersionContents.SetText("1.12.2")
	downloadPathContents.SetText("./")

	// Start button
	buttonstart := ui.NewButton("Start")
	buttonstart.OnClicked(func(*ui.Button) {
		// Set values
		setFromGUI(instancePathContents.Text(), gameVersionContents.Text(),
			downloadPathContents.Text(), exportManifestPath.Text())
		ui.Quit()
	})

	vbox.Append(buttonstart, false)

	// Booleans begin
	groupBools := ui.NewGroup("Booleans")
	groupBools.SetMargined(true)
	vbox.Append(groupBools, false)

	boolsGrid := ui.NewGrid()
	boolsGrid.SetPadded(true)
	groupBools.SetChild(boolsGrid)

	exportNewManifestBox := ui.NewCheckbox("Export manifest.json")
	boolsGrid.Append(exportNewManifestBox, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	exportOldManifestBox := ui.NewCheckbox("Export old.json")
	boolsGrid.Append(exportOldManifestBox, 0, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	silentModeBox := ui.NewCheckbox("Silent mode")
	boolsGrid.Append(silentModeBox, 0, 2, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	// Booleans end

	// Entries begin
	groupEntries := ui.NewGroup("Entries")
	groupEntries.SetMargined(true)
	vbox.Append(groupEntries, false)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	groupEntries.SetChild(entryForm)

	entryForm.Append("Instance folder path", instancePathContents, false)
	entryForm.Append("Game version", gameVersionContents, false)
	entryForm.Append("Download folder path", downloadPathContents, false)

	exportJsonButton := ui.NewButton("Open export.json file")
	exportJsonButton.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		exportManifestPath.SetText(filename)
	})

	entryForm.Append("Select export.json file", exportJsonButton, false)
	entryForm.Append("export.json path", exportManifestPath, false)
	// Entries end

	return vbox

}

func setupUI() {
	// Create the window with a title, 960x540 (16:9), with a menu bar (not sure what that is)
	mainwin = ui.NewWindow("CMPU (Curse Modpack Utilities)", 960, 540, true)

	// User closing
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	// OS closing
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	// Literally everything
	mainwin.SetChild(makeBasicControlsPage())
	mainwin.SetMargined(true)

	// Display it
	mainwin.Show()
}

func launchGUI() {

	err := ui.Main(setupUI)
	if err != nil {
		log.Fatal("Cannot create GUI.")
	}

}
