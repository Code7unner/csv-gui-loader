package gui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func Start() {
	// Инициализируем GTK.
	gtk.Init(nil)

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("main.glade")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("main_window")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//TODO: Import csv-file
	onCSVImportButtonClicked(b)

	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
}

func onCSVImportButtonClicked(b *gtk.Builder) {
	object, err := b.GetObject("file_choose_button")
	if err != nil {
		log.Fatal("Error:", err)
	}

	button := object.(*gtk.FileChooserButton)
	button.Connect("file-set", func(file *gtk.FileChooserButton) {
		println(file.GetFilename())
	})
}
