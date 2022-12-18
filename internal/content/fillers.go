package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type Calc struct {
	Equation string
	Output   *widget.Label
	Buttons  map[string]*widget.Button
	Window   fyne.Window
}

func NewCalculator() *Calc {
	return &Calc{
		Buttons: make(map[string]*widget.Button, 32),
	}
}

func (c *Calc) display(newText string) {
	c.Equation = newText
	c.Output.SetText(newText)
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
		c.digit(number)
	})
}

func (c *Calc) charButton(char rune) *widget.Button {
	return c.addButton(string(char), func() {
		c.character(char)
	})
}

func (c *Calc) stringButton(s string) *widget.Button {
	return c.addButton(s, func() {
		c.string(s)
	})
}

func (c *Calc) LoadUI(a fyne.App) {
	c.Output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.Output.TextStyle.Monospace = true

	equals := c.addButton("=", c.evaluate)
	equals.Importance = widget.HighImportance

	c.Window = a.NewWindow("Smart Calculator (c) acristin")
	c.Window.SetContent(container.NewGridWithColumns(1,
		c.Output,
		container.NewGridWithColumns(5,
			c.addButton("C", c.clear),
			c.charButton('('),
			c.charButton(')'),
			c.charButton('/'),
			c.charButton('^')),
		container.NewGridWithColumns(5,
			c.digitButton(7),
			c.digitButton(8),
			c.digitButton(9),
			c.charButton('*'),
			c.stringButton("mod")),
		container.NewGridWithColumns(5,
			c.digitButton(4),
			c.digitButton(5),
			c.digitButton(6),
			c.charButton('-'),
			c.stringButton("--")),
		container.NewGridWithColumns(5,
			c.digitButton(1),
			c.digitButton(2),
			c.digitButton(3),
			c.charButton('+'),
			c.stringButton("++")),
		container.NewGridWithColumns(5,
			c.digitButton(0),
			c.charButton('.'),
			c.stringButton("sqrt"),
			c.charButton('x'),
			equals)),
	)

	canvas := c.Window.Canvas()
	canvas.SetOnTypedRune(c.onTypedRune)
	canvas.SetOnTypedKey(c.onTypedKey)
	canvas.AddShortcut(&fyne.ShortcutCopy{}, c.onCopyShortcut)
	canvas.AddShortcut(&fyne.ShortcutPaste{}, c.onPasteShortcut)

	// add menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		//helpMenu,
	)

	// handle ESC
	c.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			a.Quit()
		}
	})
	c.Window.SetMainMenu(mainMenu)
	c.Window.Resize(fyne.NewSize(500, 300))
	c.Window.Show()

}
