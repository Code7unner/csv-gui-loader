package gui

import (
	csv_loader "../csv-loader"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func Start() {
	// Init GTK.
	gtk.Init(nil)

	// Create builder
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Import Glade-file
	err = b.AddFromFile("main.glade")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Get object by id
	obj, err := b.GetObject("main_window")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Destroy window when app closing or stopped
	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Buttons signals
	err = onCSVImportButtonClicked(b)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Display all windows
	win.ShowAll()

	// We execute the main GTK cycle (for rendering).
	// He will stop when gtk.MainQuit() is done.
	gtk.Main()
}

func onCSVImportButtonClicked(b *gtk.Builder) error {
	object, err := b.GetObject("file_choose_button")
	if err != nil {
		return err
	}

	button := object.(*gtk.FileChooserButton)
	button.Connect("file-set", func(file *gtk.FileChooserButton) {
		parsedCsvFile, _ := csv_loader.ParseCSV(file.GetFilename())

		err = addCSVFields(b, parsedCsvFile)
	})

	if err != nil {
		return err
	}

	return nil
}

func addCSVFields(b *gtk.Builder, parsedCsvFile []*csv_loader.CSV) error {
	for _, object := range parsedCsvFile {
		// TODO: create rows from parsed csv
	}
	return nil
}