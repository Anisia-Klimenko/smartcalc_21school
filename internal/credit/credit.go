package credit

import (
	"calc/internal/math/credit"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

type Credit struct {
	Sum       string
	Term      string
	Rate      string
	Monthly   *widget.Label
	Overpay   *widget.Label
	Total     *widget.Label
	IsAnnuity bool
	Entries   map[string]*widget.Entry
	Err       error
	Message   *widget.Label
}

func ShowCredit(a fyne.App) {
	log.Println("credit: opened")
	var c Credit
	c.Message = widget.NewLabel("")
	c.addResultLabels()
	// Create window and set content from file
	w4 := a.NewWindow("Credit calculator (c) acristin")
	w4.SetContent(container.NewGridWithColumns(1,
		container.NewGridWithColumns(4,
			c.addLabel("Sum"),
			c.addEntry("50000", func(s string) {
				c.Sum = s
			}),
			c.addLabel("Monthly"),
			c.Monthly,
		),
		container.NewGridWithColumns(4,
			c.addLabel("Term"),
			c.addEntry("12", func(s string) {
				c.Term = s
			}),
			c.addLabel("Overpay"),
			c.Overpay,
		),
		container.NewGridWithColumns(4,
			c.addLabel("Rate"),
			c.addEntry("4.5", func(s string) {
				c.Rate = s
			}),
			c.addLabel("Total"),
			c.Total,
		),
		container.NewGridWithColumns(4,
			c.addLabel("Type"),
			c.addRadio(),
			c.Message,
			c.addCalcButton("Calculate"),
		),
	))

	// Handle shortcuts
	w4.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			// Close window in case ESC or W was pressed
			w4.Close()
			log.Println("credit: closed")
		} else if keyEvent.Name == fyne.KeyReturn || keyEvent.Name == fyne.KeyEnter {
			// Evaluate equation in case Return or Enter was pressed
			c.Monthly.Text, c.Overpay.Text, c.Total.Text, c.Err = credit.Calculate(c.Sum, c.Term, c.Rate, c.IsAnnuity)
			if c.Err != nil {
				c.Message.SetText("error")
			} else {
				c.Message.SetText("")
				c.Monthly.SetText(c.Monthly.Text)
				c.Overpay.SetText(c.Overpay.Text)
				c.Total.SetText(c.Total.Text)
			}
		}
	})

	// Show window
	w4.Resize(fyne.NewSize(500, 300))
	w4.Show()
}
