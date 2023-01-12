package content

import (
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// addButton creates a new button widget with the set label and tap handler.
// Also adds new button to c.Buttons[]
func (c *Calc) addButton(label string, handler func()) *widget.Button {
	button := widget.NewButton(label, handler)
	c.Buttons[label] = button
	return button
}

// digitButton returns new button with the label set as number
func (c *Calc) digitButton(number int) *widget.Button {
	str := strconv.Itoa(number)
	return c.addButton(str, func() {
		// Clear output in case of new expression
		if c.ifEqualPressed || c.Output.Text == "error" {
			c.clear()
			c.ifEqualPressed = false
		}
		c.digit(number)
	})
}

// charButton returns new button with the label set as char
func (c *Calc) charButton(char rune) *widget.Button {
	return c.addButton(string(char), func() {
		// Clear output in case of new expression
		if c.ifEqualPressed || c.Output.Text == "error" {
			c.clear()
			c.ifEqualPressed = false
		}
		c.character(char)
	})
}

// stringButton returns new button with the label set as s
func (c *Calc) stringButton(s string) *widget.Button {
	// mod is a separate case because it works like an operator
	if s == "mod" {
		return c.addButton(s, func() {
			// Clear output in case of new expression
			if c.ifEqualPressed || c.Output.Text == "error" {
				c.clear()
				c.ifEqualPressed = false
			}
			c.string(s)
		})
	}
	// Add "(" to output if function button is pressed
	return c.addButton(s, func() {
		// Clear output in case of new expression
		if c.ifEqualPressed || c.Output.Text == "error" {
			c.clear()
			c.ifEqualPressed = false
		}
		c.string(s + "(")
	})
}
