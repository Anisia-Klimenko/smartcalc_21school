package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type Calc struct {
	Equation       string
	XValue         float64
	Output         *widget.Label
	Buttons        map[string]*widget.Button
	Window         fyne.Window
	ifEqualPressed bool
}

func NewCalculator() *Calc {
	return &Calc{
		Buttons: make(map[string]*widget.Button, 32),
	}
}

func (c *Calc) display(newText string) {
	c.Equation = newText
	if len(newText) <= 255 {
		c.Output.SetText(newText)
	}
}

func (c *Calc) character(char rune) {
	c.display(c.Equation + string(char))
}

func (c *Calc) digit(d int) {
	c.character(rune(d) + '0')
}

func (c *Calc) string(s string) {
	c.display(c.Equation + s)
}

func (c *Calc) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.Buttons[text] = button
	return button
}

func (c *Calc) digitButton(number int) *widget.Button {
	str := strconv.Itoa(number)
	return c.addButton(str, func() {
		if c.ifEqualPressed {
			c.clear()
			c.ifEqualPressed = false
		}
		c.digit(number)
	})
}

func (c *Calc) charButton(char rune) *widget.Button {
	return c.addButton(string(char), func() {
		if c.ifEqualPressed {
			c.clear()
			c.ifEqualPressed = false
		}
		c.character(char)
	})
}

func (c *Calc) stringButton(s string) *widget.Button {
	if s == "mod" {
		return c.addButton(s, func() {
			if c.ifEqualPressed {
				c.clear()
				c.ifEqualPressed = false
			}
			c.string(s)
		})
	}
	return c.addButton(s, func() {
		if c.ifEqualPressed {
			c.clear()
			c.ifEqualPressed = false
		}
		c.string(s + "(")
	})
}

func (c *Calc) LoadUI(a fyne.App) {
	var err error
	c.Output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.Output.TextStyle.Monospace = true

	equals := c.addButton("=", c.evaluate)
	equals.Importance = widget.HighImportance

	xEntry := widget.NewEntry()
	xEntry.TextStyle.Monospace = true
	xEntry.OnChanged = func(s string) {
		c.XValue, err = strconv.ParseFloat(s, 64)
		if err != nil {
			c.display("error")
		}
	}

	c.Window = a.NewWindow("Smart Calculator (c) acristin")
	c.Window.SetContent(container.NewGridWithColumns(1,
		container.NewHScroll(c.Output),
		container.NewGridWithColumns(7,
			c.addButton("C", c.clear),
			c.charButton('('),
			c.charButton(')'),
			c.charButton('/'),
			c.charButton('x'),
			widget.NewLabelWithStyle("x = ", fyne.TextAlignCenter, fyne.TextStyle{Monospace: true}),
			xEntry),
		container.NewGridWithColumns(7,
			c.digitButton(7),
			c.digitButton(8),
			c.digitButton(9),
			c.charButton('*'),
			c.stringButton("mod"),
			c.stringButton("sin"),
			c.stringButton("asin")),
		container.NewGridWithColumns(7,
			c.digitButton(4),
			c.digitButton(5),
			c.digitButton(6),
			c.charButton('-'),
			c.stringButton("ln"),
			c.stringButton("cos"),
			c.stringButton("acos")),
		container.NewGridWithColumns(7,
			c.digitButton(1),
			c.digitButton(2),
			c.digitButton(3),
			c.charButton('+'),
			c.stringButton("log"),
			c.stringButton("tan"),
			c.stringButton("atan")),
		container.NewGridWithColumns(7,
			c.digitButton(0),
			c.charButton('.'),
			c.stringButton("sqrt"),
			c.charButton('^'),
			equals),
	))

	canvas := c.Window.Canvas()
	canvas.SetOnTypedRune(c.onTypedRune)
	canvas.AddShortcut(&fyne.ShortcutCopy{}, c.onCopyShortcut)
	canvas.AddShortcut(&fyne.ShortcutPaste{}, c.onPasteShortcut)

	// add menu
	fileMenu := fyne.NewMenu("Calculator Menu",
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		//helpMenu,
	)
	c.Window.SetMainMenu(mainMenu)

	// handle ESC, Return, Enter, BackSpace
	c.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			a.Quit()
		} else if keyEvent.Name == fyne.KeyReturn || keyEvent.Name == fyne.KeyEnter {
			c.evaluate()
		} else if keyEvent.Name == fyne.KeyBackspace {
			c.backspace()
		}
	})
	c.Window.Resize(fyne.NewSize(500, 300))
	c.Window.Show()

}
