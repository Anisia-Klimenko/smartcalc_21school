package credit

import (
	"calc/internal/math/credit"
	"fyne.io/fyne/v2/widget"
	"log"
)

// addEntry creates a new entry widget with the set label and onchange handler.
func (c *Credit) addEntry(label string, handler func(string)) *widget.Entry {
	entry := widget.NewEntry()
	entry.TextStyle.Monospace = true
	entry.OnChanged = handler
	entry.PlaceHolder = label
	return entry
}

// addResultLabels creates result labels and sets style
func (c *Credit) addResultLabels() {
	c.Monthly = widget.NewLabel("")
	c.Monthly.TextStyle.Bold = true

	c.Overpay = widget.NewLabel("")
	c.Overpay.TextStyle.Bold = true

	c.Total = widget.NewLabel("")
	c.Total.TextStyle.Bold = true
}

// addCalcButton creates "Calculate"-button and sets handler
func (c *Credit) addCalcButton(label string) *widget.Button {
	button := widget.NewButton(label, func() {
		// Default value of radio button is set to false, this value is set to true only if radio button was set
		if !c.IsAnnuitySet {
			c.Message.SetText("error")
			log.Println("type of payment is not set")
			return
		}

		// Calculate the results
		c.Monthly.Text, c.Overpay.Text, c.Total.Text, c.Err = credit.Calculate(c.Sum, c.Term, c.Rate, c.IsAnnuity)

		if c.Err != nil {
			// Set error message label
			c.Message.SetText("error")
			log.Println(c.Err)
		} else {
			// Set results to message and result labels
			c.Message.SetText("")
			c.Monthly.SetText(c.Monthly.Text)
			c.Overpay.SetText(c.Overpay.Text)
			c.Total.SetText(c.Total.Text)
		}
	})

	// Set button style
	button.Importance = widget.HighImportance

	return button
}

// addRadio creates radio widget and sets its handler to default
func (c *Credit) addRadio() *widget.RadioGroup {
	// Options of radio widget
	options := []string{"Annuity", "Differentiated"}

	radio := widget.NewRadioGroup(options, func(s string) {
		// If radio widget value was set
		if s == "Annuity" {
			c.IsAnnuity = true
			c.IsAnnuitySet = true
		} else if s == "Differentiated" {
			c.IsAnnuity = false
			c.IsAnnuitySet = true
		}
		// Otherwise flag field IsAnnuitySet remains false,
		// so calculation will stop on validation fields
	})

	// Set radio parameters
	radio.Horizontal = false
	radio.Required = true

	return radio
}

// addLabel creates new label widget with the set label content
func (c *Credit) addLabel(label string) *widget.Label {
	return widget.NewLabel(label)
}
