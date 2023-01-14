package history

import (
	"calc/internal/file"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"strings"
)

// Flag that helps wait until equation from history is chosen
// or history window is closed
var wasShown = false

// Paths to files
var path = "../assets/log.txt"
var itemPath = "../assets/item.txt"

func UpdatePaths(newPath string) {
	path = newPath + "/assets/log.txt"
	itemPath = newPath + "/assets/item.txt"
}

// ClearHistory clears history file
func ClearHistory() {
	err := file.Clear(path)
	if err != nil {
		log.Println("history: clear: file corrupted")
	}
	log.Println("history: cleaned")
}

// GetHistory returns content of history file
func GetHistory() string {
	return file.Content(path)
}

// UpdateHistory adds content to history file
func UpdateHistory(content string) {
	err := file.Update(path, content)
	if err != nil {
		log.Println("history: update: file corrupted")
	}
}

// GetHistoryItem reads chosen equation from file
func GetHistoryItem() string {
	// History window's not closed yet, so equation's not chosen too
	for !wasShown {
	}
	result, _ := os.ReadFile(itemPath)
	err := file.Clear(itemPath)
	if err != nil {
		log.Println("history: clear: file corrupted")
		return ""
	}
	res := strings.Split(string(result), "=")[0]
	wasShown = false
	return res
}

// saveHistoryItem saves chosen equation to file
func saveHistoryItem(equation string) {
	err := file.Rewrite(itemPath, equation)
	if err != nil {
		log.Println("history: rewrite: file corrupted")
		return
	}
}

// ShowHistory opens window with operation history
func ShowHistory(a fyne.App) {
	log.Println("history: opened")
	w2 := a.NewWindow("History")
	var btns []fyne.CanvasObject

	// Read history from file
	historyFile := GetHistory()

	// Create button for every operation from file
	for _, line := range strings.Split(strings.TrimSuffix(historyFile, "\n"), "\n") {
		line := line
		if len(line) != 0 {
			btns = append(btns, widget.NewButton(line, func() {
				log.Println("operation", line, "was chosen from history")
				saveHistoryItem(line)
				wasShown = true
				w2.Close()
			}))
		} else {
			// Show message for empty history
			btns = append(btns, widget.NewLabel("Empty history"))
		}
	}

	w2.SetContent(container.NewGridWithColumns(1,
		container.NewGridWithRows(1,
			container.NewScroll(container.NewGridWithColumns(1,
				container.NewGridWithColumns(1, btns...),
			)),
		),
	))

	// Handle shortcuts
	w2.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			// Close window in case ESC or W was pressed
			wasShown = true
			w2.Close()
			log.Println("history: closed")
		} else if keyEvent.Name == fyne.KeyBackspace {
			// Clear history in case BackSpace was pressed
			ClearHistory()
			wasShown = true
			w2.Close()
		}
	})

	// Show window
	w2.Resize(fyne.NewSize(500, 200))
	w2.Show()
}
