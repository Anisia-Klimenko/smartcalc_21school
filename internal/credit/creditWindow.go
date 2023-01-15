package credit

import (
	"calc/internal/math/credit"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

type Credit struct {
	Sum          string
	Term         string
	Rate         string
	Monthly      *widget.Label
	Overpay      *widget.Label
	Total        *widget.Label
	IsAnnuity    bool
	IsAnnuitySet bool
	Err          error
	Message      *widget.Label
}

// ShowCredit opens window with credit calculator
func ShowCredit(a fyne.App) {
	log.Println("credit: opened")

	// Init empty labels
	var c Credit
	c.Message = widget.NewLabel("")
	c.addResultLabels()

	// Create window and set content
	w4 := a.NewWindow("Credit calculator (c) acristin")
	w4.SetContent(container.NewGridWithColumns(1,
		container.NewGridWithColumns(4,
			c.addLabel("Total loan amount"),
			c.addEntry("50000", func(s string) { c.Sum = s }),
			c.addLabel("Monthly payment"),
			c.Monthly,
		),
		container.NewGridWithColumns(4,
			c.addLabel("Term"),
			c.addEntry("12", func(s string) { c.Term = s }),
			c.addLabel("Loan overpayment"),
			c.Overpay,
		),
		container.NewGridWithColumns(4,
			c.addLabel("Interest rate"),
			c.addEntry("4.5", func(s string) { c.Rate = s }),
			c.addLabel("Total payment"),
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
			log.Println("credit: parameters: sum =", c.Sum, ", term =", c.Term,
				", rate =", c.Rate, ", is type annuity =", c.IsAnnuity)
			c.Monthly.Text, c.Overpay.Text, c.Total.Text, c.Err = credit.Calculate(c.Sum, c.Term, c.Rate, c.IsAnnuity)
			log.Println("credit: calculated: monthly =", c.Monthly.Text, ", overpay =",
				c.Overpay.Text, ", total =", c.Total.Text)
			if c.Err != nil {
				// Update message label
				c.Message.SetText("error")
			} else {
				// Update message and content labels
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
