package main

import (
	"calc/internal/content"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func main() {
	a := app.New()
	w := a.NewWindow("Smart Calculator (c) acristin")
	content.SetContent(&w, &a)
	w.SetContent(widget.NewLabel("Hello World!"))
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		//helpMenu,
	)
	w.SetMainMenu(mainMenu)

	var resource, _ = fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500})

	randomBtn := widget.NewButton("Random", func() {
		resource, _ := fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
		gopherImg.Resource = resource

		//Redrawn the image with the new path
		gopherImg.Refresh()
	})
	randomBtn.Importance = widget.HighImportance

	box := container.NewVBox(
		gopherImg,
		randomBtn,
	)

	// Display our content
	w.SetContent(box)

	// ESC
	w.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			a.Quit()
		}
	})
	w.ShowAndRun()
}
