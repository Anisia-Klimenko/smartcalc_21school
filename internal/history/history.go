package history

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"strings"
)

var wasShown = false

func ClearHistory() {
	f, err := os.OpenFile("../assets/log.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	f.Truncate(0)
	f.Seek(0, 0)
}

func GetHistory() string {
	result, _ := os.ReadFile("../assets/log.txt")
	return string(result)
}

func UpdateHistory(result string) {
	f, err := os.OpenFile("../assets/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(result + "\n"); err != nil {
		log.Println(err)
	}
}

func GetHistoryItem() string {
	for !wasShown {
	}
	result, _ := os.ReadFile("../assets/item.txt")
	f, err := os.OpenFile("../assets/item.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	f.Truncate(0)
	f.Seek(0, 0)
	res := strings.Split(string(result), "=")[0]
	log.Println("getting", res, "...")
	wasShown = false
	return res
}

func saveHistoryItem(result string) {
	f, err := os.OpenFile("../assets/item.txt", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(result + "\n"); err != nil {
		log.Println(err)
	}
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
