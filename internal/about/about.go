package about

import (
	"calc/internal/file"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

var path = "../assets/about.txt"

// UpdatePath updates path if app was installed as desktop app
func UpdatePath(newPath string) {
	path = newPath + "/assets/about.txt"
}

// getAbout reads file with reference and return content in string format
func getAbout() string {
	about := file.Content(path)
	// In case content is missing
	if len(about) == 0 {
		about = "Content is missing " + path
	}
	return about
}

// ShowAbout opens window with reference
func ShowAbout(a fyne.App) {
	log.Println("about: opened")

	// Create window and set content from file
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
