package about

import (
	"calc/internal/file"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

// getAbout reads file with reference and return content in string format
func getAbout() string {
	about := file.Content("../assets/about.txt")
	// In case content is missing
	if len(about) == 0 {
		return "Content is missing"
	}
	return about
}

// ShowAbout opens window with reference
func ShowAbout(a fyne.App) {
	log.Println("about: opened")
	w3 := a.NewWindow("About")
	about := getAbout() // Get content from file
	w3.SetContent(container.NewGridWithColumns(1,
		container.NewVScroll(widget.NewLabel(about)),
	))

	// Handle shortcuts
	w3.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			// Close window in case ESC or W was pressed
			w3.Close()
			log.Println("about: closed")
		}
	})

	// Show window
	w3.Resize(fyne.NewSize(500, 400))
	w3.Show()
}
