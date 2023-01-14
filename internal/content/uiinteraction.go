package content

import (
	"calc/internal/history"
	"calc/internal/math"
	"fyne.io/fyne/v2"
	"log"
)

// backspace removes last typed symbol
func (c *Calc) backspace() {
	if len(c.Equation) == 0 || c.Equation == "error" {
		// Empty equation
		return
	} else if c.Equation == "error" {
		// Clear output if error was typed
		c.clear()
		return
	}

	c.display(c.Equation[:len(c.Equation)-1])
}

// onTypedRune runs button handler if it is present in calculator object button slice
func (c *Calc) onTypedRune(r rune) {
	if r == 'c' {
		r = 'C' // The button is using an uppercase C.
	}
	if r == 'E' {
		r = 'e' // The button is using a lowercase e.
	}

	if button, ok := c.Buttons[string(r)]; ok {
		// Launch button handler
		button.OnTapped()
	}
}

// display sets calculator output to newText
func (c *Calc) display(newText string) {
	// Change equation value, that will be calculated later
	c.Equation = newText
	if len(newText) <= 255 {
		// Set output of calculator, if addition of new text
		// is in 255-symbol limit
		c.Output.SetText(newText)
	}
}

// character adds char to displayed output
func (c *Calc) character(char rune) {
	c.display(c.Equation + string(char))
}

// digit adds number to displayed output
func (c *Calc) digit(number int) {
	c.character(rune(number) + '0')
}

// string adds str to displayed output
func (c *Calc) string(str string) {
	c.display(c.Equation + str)
}

// clear clears output
func (c *Calc) clear() {
	c.display("")
}

// evaluate displays equation result and saves expression to
// history file
func (c *Calc) evaluate() {
	// Change flag value
	c.ifEqualPressed = true

	// Calculate result
	result := math.Calculate(c.Output.Text, c.XValue)

	// Handle history
	var historyText = c.Output.Text + "="
	historyText += result
	history.UpdateHistory(historyText)

	// Display result
	c.display(result)
	log.Println(historyText)
}

// onPasteShortcut handles cmd+P shortcut
func (c *Calc) onPasteShortcut(shortcut fyne.Shortcut) {
	// Get content from clipboard
	content := shortcut.(*fyne.ShortcutPaste).Clipboard.Content()
	// Display it
	c.display(c.Equation + content)
}

// onCopyShortcut handles cmd+C shortcut
func (c *Calc) onCopyShortcut(shortcut fyne.Shortcut) {
	shortcut.(*fyne.ShortcutCopy).Clipboard.SetContent(c.Equation)
}
