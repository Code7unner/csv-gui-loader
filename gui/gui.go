package gui

import (
	csv_loader "../csv-loader"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
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
	onCSVImportButtonClicked(b)

	// Display all windows
	win.ShowAll()

	// We execute the main GTK cycle (for rendering).
	// He will stop when gtk.MainQuit() is done.
	gtk.Main()
}

func onCSVImportButtonClicked(b *gtk.Builder) {
	object, err := b.GetObject("file_choose_button")
	if err != nil {
		log.Fatal("Error:", err)
	}

	button := object.(*gtk.FileChooserButton)
	button.Connect("file-set", func(file *gtk.FileChooserButton) {
		unparsedCsvFile, err := os.Open(file.GetFilename())
		if err != nil {
			log.Fatal("Error:", err)
		}

		reader := csv.NewReader(bufio.NewReader(unparsedCsvFile))
		parsedCsvFile := csv_loader.ParseCSV(reader)

		fmt.Println(string(parsedCsvFile))
	})
}
