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

var wasShown = false

func ClearHistory() {
	file.Clear("../assets/log.txt")
}

func GetHistory() string {
	return file.Content("../assets/log.txt")
}

func UpdateHistory(content string) {
	file.Update("../assets/log.txt", content)
}

func GetHistoryItem() string {
	for !wasShown {
	}
	result, _ := os.ReadFile("../assets/item.txt")
	file.Clear("../assets/item.txt")
	res := strings.Split(string(result), "=")[0]
	wasShown = false
	return res
}

func saveHistoryItem(content string) {
	file.Rewrite("../assets/item.txt", content)
}

func ShowHistory(a fyne.App) {
	w2 := a.NewWindow("History")
	var btns []fyne.CanvasObject
	file := GetHistory()
	for _, line := range strings.Split(strings.TrimSuffix(file, "\n"), "\n") {
		line := line
		if len(line) != 0 {
			btns = append(btns, widget.NewButton(line, func() {
				log.Println(line)
				saveHistoryItem(line)
				wasShown = true
				w2.Close()
			}))
		} else {
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
	w2.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			wasShown = true
			w2.Close()
		} else if keyEvent.Name == fyne.KeyBackspace {
			ClearHistory()
			wasShown = true
			w2.Close()
		}
	})
	w2.Resize(fyne.NewSize(500, 200))
	w2.Show()
}
