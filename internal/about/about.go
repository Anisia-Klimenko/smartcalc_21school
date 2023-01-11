package about

import (
	"calc/internal/file"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetAbout() string {
	about := file.Content("../assets/about.txt")
	if len(about) == 0 {
		return "Content is missing"
	}
	return about
}

func ShowAbout(a fyne.App) {
	w3 := a.NewWindow("About")
	about := GetAbout()
	w3.SetContent(container.NewGridWithColumns(1,
		container.NewVScroll(widget.NewLabel(about)),
	))
	w3.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			w3.Close()
		}
	})
	w3.Resize(fyne.NewSize(500, 400))
	w3.Show()
}
