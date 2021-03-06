package main

import (
	"log"
	"sync"

	"github.com/gotk3/gotk3/gtk"
)

func launchGUI() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("CMPU")

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a GtkFixed widget to show in the window.
	fixed, err := gtk.FixedNew()

	if err != nil {
		log.Fatal("Unable to create GtkFixed:", err)
	}

	/*
		Add items to the GtkFixed
	*/
	// Start button
	start, _ := gtk.ButtonNewWithLabel("Run")
	start.SetSizeRequest(800, 25)
	fixed.Put(start, 0, 0)

	/*
		Checkboxes
	*/
	exportManifestChk, _ := gtk.CheckButtonNewWithLabel("Export manifest.json")
	fixed.Put(exportManifestChk, 5, 50)

	exportOldChk, _ := gtk.CheckButtonNewWithLabel("Export old.json")
	fixed.Put(exportOldChk, 5, 75)

	silentModeChk, _ := gtk.CheckButtonNewWithLabel("Silent mode")
	fixed.Put(silentModeChk, 5, 100)

	/*
		Entries
	*/
	// Instance folder
	instanceFolderLbl, _ := gtk.LabelNew("Instance folder path: ")
	instanceFolderTv, _ := gtk.EntryNew()
	instanceFolderTv.SetText("./") // Default value
	instanceFolderTv.SetSizeRequest(295, 10)
	instanceFolderBtn, _ := gtk.FileChooserButtonNew("Instance folder path", gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER)
	fixed.Put(instanceFolderLbl, 5, 133)
	fixed.Put(instanceFolderTv, 150, 125)
	fixed.Put(instanceFolderBtn, 450, 125)

	// Game version
	gameVersionLbl, _ := gtk.LabelNew("Game version: ")
	gameVersionTv, _ := gtk.EntryNew()
	gameVersionTv.SetText("1.12.2") // Default value
	gameVersionTv.SetSizeRequest(295, 10)
	fixed.Put(gameVersionLbl, 5, 183)
	fixed.Put(gameVersionTv, 150, 175)

	// Release type
	releaseTypeLbl, _ := gtk.LabelNew("Release type: ")
	releaseTypeComboBox, _ := gtk.ComboBoxTextNew()
	releaseTypeComboBox.AppendText("stable")
	releaseTypeComboBox.AppendText("latest")
	releaseTypeComboBox.SetActive(0) // Default value
	releaseTypeComboBox.SetSizeRequest(295, 10)
	fixed.Put(releaseTypeLbl, 5, 233)
	fixed.Put(releaseTypeComboBox, 150, 225)

	// Download folder
	downloadFolderLbl, _ := gtk.LabelNew("Download folder path: ")
	downloadFolderTv, _ := gtk.EntryNew()
	downloadFolderTv.SetText("./") // Default value
	downloadFolderTv.SetSizeRequest(295, 10)
	downloadFolderBtn, _ := gtk.FileChooserButtonNew("Download folder path", gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER)
	fixed.Put(downloadFolderLbl, 5, 283)
	fixed.Put(downloadFolderTv, 150, 275)
	fixed.Put(downloadFolderBtn, 450, 275)

	// export.json file
	exportJSONLbl, _ := gtk.LabelNew("export.json path: ")
	exportJSONTv, _ := gtk.EntryNew()
	exportJSONTv.SetText("./") // Default value
	exportJSONTv.SetSizeRequest(295, 10)
	exportJSONBtn, _ := gtk.FileChooserButtonNew("export.json file path", gtk.FILE_CHOOSER_ACTION_OPEN)
	fixed.Put(exportJSONLbl, 5, 333)
	fixed.Put(exportJSONTv, 150, 325)
	fixed.Put(exportJSONBtn, 450, 325)

	// Add the GtkFixed to the window.
	win.Add(fixed)

	/*
	   Start execution
	   of program button
	   and actually do it...
	*/
	var wg sync.WaitGroup
	start.Connect("clicked", func() {
		// Set a lot of the values
		*exportNewManifest = exportManifestChk.GetActive()
		*exportOldManifest = exportOldChk.GetActive()
		*silentMode = silentModeChk.GetActive()

		*instancePath, _ = instanceFolderTv.GetText()
		gameVersionText, _ := gameVersionTv.GetText()
		*gameVersion = gameVersionText
		*releaseType = releaseTypeComboBox.GetActiveText()
		*downloadPath, _ = downloadFolderTv.GetText()
		*exportManifestPath, _ = exportJSONTv.GetText()

		// Run stuff in separate `thread`
		// Now begin
		log.Printf("Starting execution of program...")
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			readInstancePath()
			useArgs()
			checkUpdates(oldMap, newMap)
		}(&wg)
		wg.Wait() // Lock until complete
	})

	/*
		Event handlers
	*/
	exportManifestChk.Connect("toggled", func() {
		log.Printf("Export manifest.json: %t", exportManifestChk.GetActive())
	})

	exportOldChk.Connect("toggled", func() {
		log.Printf("Export old.json: %t", exportOldChk.GetActive())
	})

	silentModeChk.Connect("toggled", func() {
		log.Printf("Silent mode: %t", silentModeChk.GetActive())
	})

	instanceFolderBtn.Connect("file-set", func() {
		folderPath := instanceFolderBtn.GetFilename()
		instanceFolderTv.SetText(folderPath)
		log.Printf("Instance folder path set to: %s", folderPath)
	})

	downloadFolderBtn.Connect("file-set", func() {
		folderPath := downloadFolderBtn.GetFilename()
		downloadFolderTv.SetText(folderPath)
		log.Printf("Download folder path set to: %s", folderPath)
	})

	exportJSONBtn.Connect("file-set", func() {
		filePath := exportJSONBtn.GetFilename()
		exportJSONTv.SetText(filePath)
		log.Printf("export.json file path set to: %v", filePath)
	})

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
