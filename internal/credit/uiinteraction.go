package credit

import (
	"calc/internal/math/credit"
	"fyne.io/fyne/v2/widget"
	"log"
)

func (c *Credit) addEntry(label string, handler func(string)) *widget.Entry {
	entry := widget.NewEntry()
	entry.TextStyle.Monospace = true
	entry.OnChanged = handler
	entry.PlaceHolder = label
	return entry
}

func (c *Credit) addResultLabels() {
	c.Monthly = widget.NewLabel("")
	c.Monthly.TextStyle.Bold = true

	c.Overpay = widget.NewLabel("")
	c.Overpay.TextStyle.Bold = true

	c.Total = widget.NewLabel("")
	c.Total.TextStyle.Bold = true
}

func (c *Credit) addCalcButton(label string) *widget.Button {
	button := widget.NewButton(label, func() {
		//c.Monthly.Text, c.Overpay.Text, c.Total.Text, c.Err = credit.Calculate("50000", "12", "12", true)
		if !c.IsAnnuitySet {
			c.Message.SetText("error")
			log.Println("type of payment is not set")
			return
		}
		c.Monthly.Text, c.Overpay.Text, c.Total.Text, c.Err = credit.Calculate(c.Sum, c.Term, c.Rate, c.IsAnnuity)
		if c.Err != nil {
			c.Message.SetText("error")
			log.Println(c.Err)
		} else {
			c.Message.SetText("")
			c.Monthly.SetText(c.Monthly.Text)
			c.Overpay.SetText(c.Overpay.Text)
			c.Total.SetText(c.Total.Text)
		}
	})
	button.Importance = widget.HighImportance
	return button
}

func (c *Credit) addRadio() *widget.RadioGroup {
	options := []string{"Annuity", "Differentiated"}
	radio := widget.NewRadioGroup(options, func(s string) {
		if s == "Annuity" {
			c.IsAnnuity = true
			c.IsAnnuitySet = true
		} else if s == "Differentiated" {
			c.IsAnnuity = false
			c.IsAnnuitySet = true
		}
	})
	radio.Horizontal = false
	radio.Required = true
	return radio
}

func (c *Credit) addLabel(label string) *widget.Label {
	return widget.NewLabel(label)
}
