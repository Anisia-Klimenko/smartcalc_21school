package content

import (
	"calc/internal/history"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
)

type Calc struct {
	Equation       string
	XValue         float64
	Borders        []float64
	Output         *widget.Label
	Buttons        map[string]*widget.Button
	Entries        map[string]*widget.Entry
	Window         fyne.Window
	ifEqualPressed bool
	App            fyne.App
}

func NewCalculator(a fyne.App) *Calc {
	return &Calc{
		App:     a,
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

func (c *Calc) changeXValue(s string) {
	var err error
	c.XValue, err = strconv.ParseFloat(s, 64)
	if err != nil {
		c.display("error")
	}
}

func (c *Calc) addEntry(text string, action func(string)) *widget.Entry {
	entry := widget.NewEntry()
	entry.TextStyle.Monospace = true
	entry.OnChanged = action
	entry.PlaceHolder = text
	log.Println(text, len(text))
	return entry
}

func (c *Calc) CopyEquation(text string) {
	c.display(text)
}

func (c *Calc) LoadUI(a fyne.App) {
	c.Output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.Output.TextStyle.Monospace = true

	equals := c.addButton("=", c.evaluate)
	equals.Importance = widget.HighImportance

	xEntry := widget.NewEntry()
	xEntry.TextStyle.Monospace = true
	xEntry.OnChanged = c.changeXValue

	c.Window = a.NewWindow("Smart Calculator (c) acristin")
	c.Window.SetContent(container.NewGridWithColumns(1,
		container.NewHScroll(c.Output),
		container.NewGridWithColumns(8,
			c.addButton("C", c.clear),
			c.charButton('('),
			c.charButton(')'),
			c.charButton('/'),
			c.charButton('x'),
			c.addEntry("x", c.changeXValue)),
		container.NewGridWithColumns(8,
			c.digitButton(7),
			c.digitButton(8),
			c.digitButton(9),
			c.charButton('*'),
			c.stringButton("mod"),
			c.stringButton("sin"),
			c.stringButton("asin"),
			c.addEntry("x min", c.changeXValue)),
		container.NewGridWithColumns(8,
			c.digitButton(4),
			c.digitButton(5),
			c.digitButton(6),
			c.charButton('-'),
			c.stringButton("ln"),
			c.stringButton("cos"),
			c.stringButton("acos"),
			c.addEntry("x max", c.changeXValue)),
		container.NewGridWithColumns(8,
			c.digitButton(1),
			c.digitButton(2),
			c.digitButton(3),
			c.charButton('+'),
			c.stringButton("log"),
			c.stringButton("tan"),
			c.stringButton("atan"),
			c.addEntry("y min", c.changeXValue)),
		container.NewGridWithColumns(8,
			c.digitButton(0),
			c.charButton('.'),
			c.stringButton("sqrt"),
			c.charButton('^'),
			c.charButton('e'),
			equals,
			c.addButton("plot", c.generatePlot),
			c.addEntry("y max", c.changeXValue)),
	))

	canvas := c.Window.Canvas()
	canvas.SetOnTypedRune(c.onTypedRune)
	canvas.AddShortcut(&fyne.ShortcutCopy{}, c.onCopyShortcut)
	canvas.AddShortcut(&fyne.ShortcutPaste{}, c.onPasteShortcut)

	// handle ESC, Return, Enter, BackSpace
	c.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			a.Quit()
		} else if keyEvent.Name == fyne.KeyReturn || keyEvent.Name == fyne.KeyEnter {
			c.evaluate()
		} else if keyEvent.Name == fyne.KeyBackspace {
			c.backspace()
		} else if keyEvent.Name == "H" {
			log.Println("history")
			history.ShowHistory(a)
			operation := history.GetHistoryItem()
			if len(operation) != 0 {
				c.display(operation)
			}
		}
	})

	c.Window.Resize(fyne.NewSize(600, 300))

	// add menu
	fileMenu := fyne.NewMenu("Calculator Menu",
		fyne.NewMenuItem("About", func() {

		}),
		fyne.NewMenuItem("Quit", func() { a.Quit() }),
	)
	historyMenu := fyne.NewMenu("History",
		fyne.NewMenuItem("Show", func() {
			history.ShowHistory(a)
		}),
		fyne.NewMenuItem("Clear", history.ClearHistory),
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		historyMenu,
	)
	c.Window.SetMainMenu(mainMenu)
	c.Window.Show()
}
