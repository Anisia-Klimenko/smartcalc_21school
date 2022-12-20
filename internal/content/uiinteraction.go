package content

import (
	"fyne.io/fyne/v2"
	"strconv"
)

func (c *Calc) backspace() {
	if len(c.Equation) == 0 {
		return
	} else if c.Equation == "error" {
		c.clear()
		return
	}

	c.display(c.Equation[:len(c.Equation)-1])
}

func (c *Calc) onTypedRune(r rune) {
	if r == 'c' {
		r = 'C' // The button is using a capital C.
	}

	if button, ok := c.Buttons[string(r)]; ok {
		button.OnTapped()
	}
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
