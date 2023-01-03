package history

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
)

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

func SaveHistory(result string) {
	f, err := os.OpenFile("../assets/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
	content := container.NewGridWithRows(1,
		container.NewScroll(widget.NewLabel(GetHistory())),
	)
	content.Resize(fyne.Size{Height: 150})
	clear := widget.NewButton("Clear", ClearHistory)
	clear.Resize(fyne.Size{Height: 10})
	buts := container.New(
		layout.NewHBoxLayout(),
		//container.NewGridWithColumns(2,
		layout.NewSpacer(),
		//widget.NewButton("Clear", ClearHistory),
		clear,
		layout.NewSpacer(),
		widget.NewButton("Close", w2.Close),
		layout.NewSpacer())
	buts.Resize(fyne.Size{Width: 10})

	w2.SetContent(container.NewGridWithColumns(1,
		content,
		buts,
		//container.New(
		//	layout.NewHBoxLayout(),
		//	//container.NewGridWithColumns(2,
		//	widget.NewButton("Clear", ClearHistory),
		//	layout.NewSpacer(),
		//	widget.NewButton("Close", w2.Close))),
	))
	w2.Resize(fyne.NewSize(500, 200))
	w2.Show()
}
