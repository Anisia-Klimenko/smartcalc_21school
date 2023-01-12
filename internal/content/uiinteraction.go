package content

import (
	"calc/internal/history"
	"calc/internal/math"
	"fyne.io/fyne/v2"
	"log"
	"strconv"
)

// backspace pushes the typewriter carriage one position backwards
func (c *Calc) backspace() {
	if len(c.Equation) == 0 {
		return
	} else if c.Equation == "error" {
		c.clear()
		return
	}

	c.display(c.Equation[:len(c.Equation)-1])
}

// onTypedRune runs button handler if it is present in c.Buttons
func (c *Calc) onTypedRune(r rune) {
	if r == 'c' {
		r = 'C' // The button is using an uppercase C.
	}
	if r == 'E' {
		r = 'e' // The button is using a lowercase e.
	}

	if button, ok := c.Buttons[string(r)]; ok {
		button.OnTapped()
	}
}

// display sets c.Output value to newText, changes c.Equation value to newText
func (c *Calc) display(newText string) {
	c.Equation = newText
	if len(newText) <= 255 {
		c.Output.SetText(newText)
	}
}

// character adds char to displayed output
func (c *Calc) character(char rune) {
	c.display(c.Equation + string(char))
}

// digit adds d to displayed output
func (c *Calc) digit(d int) {
	c.character(rune(d) + '0')
}

// string adds s to displayed output
func (c *Calc) string(s string) {
	c.display(c.Equation + s)
}

// clear clears output
func (c *Calc) clear() {
	c.display("")
}

// evaluate displays equation result and saves expression to
// history file
func (c *Calc) evaluate() {
	var historyText = c.Output.Text + "="
	c.ifEqualPressed = true
	result := math.Calculate(c.Output.Text, c.XValue)
	historyText += result
	history.UpdateHistory(historyText)
	c.display(result)
	log.Println(historyText)
}

func (c *Calc) onPasteShortcut(shortcut fyne.Shortcut) {
	content := shortcut.(*fyne.ShortcutPaste).Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err != nil {
		return
	}

	c.display(c.Equation + content)
}

func (c *Calc) onCopyShortcut(shortcut fyne.Shortcut) {
	shortcut.(*fyne.ShortcutCopy).Clipboard.SetContent(c.Equation)
}
