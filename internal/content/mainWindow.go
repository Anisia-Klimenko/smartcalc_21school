package content

import (
	"calc/internal/about"
	"calc/internal/history"
	"calc/internal/plot"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
)

type Borders struct {
	XMin float64
	XMax float64
	YMin float64
	YMax float64
}

type Calc struct {
	Equation       string
	XValue         float64
	Border         Borders
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
		if c.ifEqualPressed || c.Output.Text == "error" {
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

func (c *Calc) changeXValue(s string, val *float64) {
	*val, _ = strconv.ParseFloat(s, 64)
}

func (c *Calc) addEntry(text string, action func(string)) *widget.Entry {
	entry := widget.NewEntry()
	entry.TextStyle.Monospace = true
	entry.OnChanged = action
	entry.PlaceHolder = text
	return entry
}

func (c *Calc) CopyEquation(text string) {
	c.display(text)
}

func isInLimits(val float64) bool {
	if val >= -1000000 && val <= 1000000 {
		return true
	}
	return false
}

func (c *Calc) checkBorders() {
	if c.Border.XMin == 0 && c.Border.XMax == 0 {
		c.Border.XMin = -10
		c.Border.XMax = 10
	}
	if c.Border.YMin == 0 && c.Border.YMax == 0 {
		c.Border.YMin = -30
		c.Border.YMax = 30
	}
	if c.Border.XMin >= c.Border.XMax || c.Border.YMin >= c.Border.YMax ||
		!isInLimits(c.Border.XMin) || !isInLimits(c.Border.XMax) ||
		!isInLimits(c.Border.YMin) || !isInLimits(c.Border.YMax) {
		c.Border.XMin = 0
		c.Border.YMin = 0
		c.Border.XMax = 0
		c.Border.YMax = 0
		c.display("non valid borders")
	}
}

func (c *Calc) LoadUI(a fyne.App) {
	log.Println("app: start ...")
	c.Output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.Output.TextStyle.Monospace = true

	equals := c.addButton("=", c.evaluate)
	equals.Importance = widget.HighImportance

	c.Window = a.NewWindow("Smart Calculator (c) acristin")
	c.Window.SetContent(container.NewGridWithColumns(1,
		container.NewHScroll(c.Output),
		container.NewGridWithColumns(8,
			c.addButton("C", c.clear),
			c.charButton('('),
			c.charButton(')'),
			c.charButton('/'),
			c.charButton('x'),
			c.addEntry("x", func(s string) { c.changeXValue(s, &c.XValue) })),
		container.NewGridWithColumns(8,
			c.digitButton(7),
			c.digitButton(8),
			c.digitButton(9),
			c.charButton('*'),
			c.stringButton("mod"),
			c.stringButton("sin"),
			c.stringButton("asin"),
			c.addEntry("x min", func(s string) { c.changeXValue(s, &c.Border.XMin) })),
		container.NewGridWithColumns(8,
			c.digitButton(4),
			c.digitButton(5),
			c.digitButton(6),
			c.charButton('-'),
			c.stringButton("ln"),
			c.stringButton("cos"),
			c.stringButton("acos"),
			c.addEntry("x max", func(s string) { c.changeXValue(s, &c.Border.XMax) })),
		container.NewGridWithColumns(8,
			c.digitButton(1),
			c.digitButton(2),
			c.digitButton(3),
			c.charButton('+'),
			c.stringButton("log"),
			c.stringButton("tan"),
			c.stringButton("atan"),
			c.addEntry("y min", func(s string) { c.changeXValue(s, &c.Border.YMin) })),
		container.NewGridWithColumns(8,
			c.digitButton(0),
			c.charButton('.'),
			c.stringButton("sqrt"),
			c.charButton('^'),
			c.charButton('e'),
			equals,
			c.addButton("plot", func() {
				if c.Equation != "error" {
					c.checkBorders()
					plot.ShowPlot(a, c.Equation, plot.Borders(c.Border))
				}
			}),
			c.addEntry("y max", func(s string) { c.changeXValue(s, &c.Border.YMax) })),
	))

	canvas := c.Window.Canvas()
	canvas.SetOnTypedRune(c.onTypedRune)
	canvas.AddShortcut(&fyne.ShortcutCopy{}, c.onCopyShortcut)
	canvas.AddShortcut(&fyne.ShortcutPaste{}, c.onPasteShortcut)

	// handle ESC, Return, Enter, BackSpace
	c.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			log.Println("app: quit ...")
			a.Quit()
		} else if keyEvent.Name == fyne.KeyReturn || keyEvent.Name == fyne.KeyEnter {
			c.evaluate()
		} else if keyEvent.Name == fyne.KeyBackspace {
			c.backspace()
		} else if keyEvent.Name == "H" {
			history.ShowHistory(a)
			operation := history.GetHistoryItem()
			if len(operation) != 0 {
				c.display(operation)
			}
		} else if keyEvent.Name == "A" {
			about.ShowAbout(a)
		} else if keyEvent.Name == "P" {
			c.checkBorders()
			plot.ShowPlot(a, c.Equation, plot.Borders(c.Border))
		}
	})

	c.Window.Resize(fyne.NewSize(600, 300))

	// add menu
	fileMenu := fyne.NewMenu("Calculator Menu",
		fyne.NewMenuItem("About", func() { about.ShowAbout(a) }),
		fyne.NewMenuItem("Quit", func() {
			log.Println("app: quit ...")
			a.Quit()
		}),
	)
	historyMenu := fyne.NewMenu("History",
		fyne.NewMenuItem("Show", func() {
			history.ShowHistory(a)
			operation := history.GetHistoryItem()
			if len(operation) != 0 {
				c.display(operation)
			}
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
