package content

import (
	"calc/internal/about"
	"calc/internal/history"
	"calc/internal/plot"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

// Borders includes borders for plot
type Borders struct {
	XMin float64
	XMax float64
	YMin float64
	YMax float64
}

// IsBorderSet indicates if borders for plot set
type IsBorderSet struct {
	XMin bool
	XMax bool
	YMin bool
	YMax bool
}

type Calc struct {
	Equation       string
	XValue         float64
	Border         Borders
	IsBorderSet    IsBorderSet
	Output         *widget.Label
	Buttons        map[string]*widget.Button
	Entries        map[string]*widget.Entry
	Window         fyne.Window
	ifEqualPressed bool
	App            fyne.App
}

// NewCalculator creates new Calc object
func NewCalculator(a fyne.App) *Calc {
	return &Calc{
		App:     a,
		Buttons: make(map[string]*widget.Button, 32),
	}
}

// LoadUI creates main window, interface and menu, calls handlers
func (c *Calc) LoadUI(a fyne.App) {
	log.Println("app: start ...")

	// Creates output and set output style
	c.Output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	c.Output.TextStyle.Monospace = true

	// Create equal button and set its importance
	equals := c.addButton("=", c.evaluate)
	equals.Importance = widget.HighImportance

	// Create main window and set content
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
			c.addEntry("x min", func(s string) { c.changeBorderValue(s, &c.Border.XMin, &c.IsBorderSet.XMin) })),
		container.NewGridWithColumns(8,
			c.digitButton(4),
			c.digitButton(5),
			c.digitButton(6),
			c.charButton('-'),
			c.stringButton("ln"),
			c.stringButton("cos"),
			c.stringButton("acos"),
			c.addEntry("x max", func(s string) { c.changeBorderValue(s, &c.Border.XMax, &c.IsBorderSet.XMax) })),
		container.NewGridWithColumns(8,
			c.digitButton(1),
			c.digitButton(2),
			c.digitButton(3),
			c.charButton('+'),
			c.stringButton("log"),
			c.stringButton("tan"),
			c.stringButton("atan"),
			c.addEntry("y min", func(s string) { c.changeBorderValue(s, &c.Border.YMin, &c.IsBorderSet.YMin) })),
		container.NewGridWithColumns(8,
			c.digitButton(0),
			c.charButton('.'),
			c.stringButton("sqrt"),
			c.charButton('^'),
			c.charButton('e'),
			equals,
			c.addButton("plot", func() {
				if c.Equation != "error" {
					if c.checkBorders() {
						plot.ShowPlot(a, c.Equation, plot.Borders(c.Border))
					}
				}
			}),
			c.addEntry("y max", func(s string) { c.changeBorderValue(s, &c.Border.YMax, &c.IsBorderSet.YMax) })),
	))

	// Handle typed symbols
	canvas := c.Window.Canvas()
	canvas.SetOnTypedRune(c.onTypedRune)

	// Handle shortcuts
	canvas.AddShortcut(&fyne.ShortcutCopy{}, c.onCopyShortcut)   // cmd + C
	canvas.AddShortcut(&fyne.ShortcutPaste{}, c.onPasteShortcut) // cmd + V
	c.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			// Quit app in case ESC or W was pressed
			log.Println("app: quit ...")
			a.Quit()
		} else if keyEvent.Name == fyne.KeyReturn || keyEvent.Name == fyne.KeyEnter {
			// Evaluate equation in case Return or Enter was pressed
			c.evaluate()
		} else if keyEvent.Name == fyne.KeyBackspace {
			// Push the typewriter carriage one position backwards in case
			// BackSpace was pressed
			c.backspace()
		} else if keyEvent.Name == "H" {
			// Open history in case H was pressed
			history.ShowHistory(a)
			// Get chosen operation
			operation := history.GetHistoryItem()
			// Display operation if it was chosen
			if len(operation) != 0 {
				c.display(operation)
			}
		} else if keyEvent.Name == "A" {
			// Open reference in case A was pressed
			about.ShowAbout(a)
		} else if keyEvent.Name == "P" {
			// Build plot in case P was pressed
			if c.checkBorders() {
				plot.ShowPlot(a, c.Equation, plot.Borders(c.Border))
			}
		}
	})

	// Add app menu
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

	// Show window
	c.Window.Resize(fyne.NewSize(600, 300))
	c.Window.Show()
}
