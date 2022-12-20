package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Plot struct {
	Equation string
	XMin     float64
	XMax     float64
	YMin     float64
	YMax     float64
	//Output   *widget.Label
	Buttons map[string]*widget.Button
	Window  fyne.Window
	//Chart          fyne.Window
	//ifEqualPressed bool
}

func NewPlot(eq string) *Plot {
	return &Plot{
		Equation: eq,
		XMin:     -1000000,
		XMax:     1000000,
		YMin:     -1000000,
		YMax:     1000000,
		Buttons:  make(map[string]*widget.Button, 1),
	}
}

func (p *Plot) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	p.Buttons[text] = button
	return button
}

func (c *Calc) generatePlot() {
	//p := NewPlot(c.Equation)
	c.Window = c.App.NewWindow("Plot (c) acristin")
}

func (p *Plot) LoadPlotWindow(a fyne.App) {
	//var err error
	////p.Output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	////p.Output.TextStyle.Monospace = true
	//
	//generate := p.addButton("=", p.evaluate)
	//generate.Importance = widget.HighImportance
	//
	//xEntry := widget.NewEntry()
	//xEntry.TextStyle.Monospace = true
	//xEntry.OnChanged = func(s string) {
	//	p.XValue, err = strconv.ParseFloat(s, 64)
	//	if err != nil {
	//		p.display("error")
	//	}
	//}
	//
	//p.Window = a.NewWindow("Smart Calculator (c) acristin")
	//p.Window.SetContent(container.NewGridWithColumns(1,
	//	container.NewHScroll(p.Output),
	//	container.NewGridWithColumns(7,
	//		p.addButton("C", p.clear),
	//		p.charButton('('),
	//		p.charButton(')'),
	//		p.charButton('/'),
	//		p.charButton('x'),
	//		widget.NewLabelWithStyle("x = ", fyne.TextAlignCenter, fyne.TextStyle{Monospace: true}),
	//		xEntry),
	//	container.NewGridWithColumns(7,
	//		p.digitButton(7),
	//		p.digitButton(8),
	//		p.digitButton(9),
	//		p.charButton('*'),
	//		p.stringButton("mod"),
	//		p.stringButton("sin"),
	//		p.stringButton("asin")),
	//	container.NewGridWithColumns(7,
	//		p.digitButton(4),
	//		p.digitButton(5),
	//		p.digitButton(6),
	//		p.charButton('-'),
	//		p.stringButton("ln"),
	//		p.stringButton("cos"),
	//		p.stringButton("acos")),
	//	container.NewGridWithColumns(7,
	//		p.digitButton(1),
	//		p.digitButton(2),
	//		p.digitButton(3),
	//		p.charButton('+'),
	//		p.stringButton("log"),
	//		p.stringButton("tan"),
	//		p.stringButton("atan")),
	//	container.NewGridWithColumns(7,
	//		p.digitButton(0),
	//		p.charButton('.'),
	//		p.stringButton("sqrt"),
	//		p.charButton('^'),
	//		p.charButton('e'),
	//		equals),
	//))
	//
	//canvas := p.Window.Canvas()
	//canvas.SetOnTypedRune(p.onTypedRune)
	//canvas.AddShortcut(&fyne.ShortcutCopy{}, p.onCopyShortcut)
	//canvas.AddShortcut(&fyne.ShortcutPaste{}, p.onPasteShortcut)
	//
	//// handle ESC, Return, Enter, BackSpace
	//p.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
	//	if keyEvent.Name == fyne.KeyEscape {
	//		a.Quit()
	//	} else if keyEvent.Name == fyne.KeyReturn || keyEvent.Name == fyne.KeyEnter {
	//		p.evaluate()
	//	} else if keyEvent.Name == fyne.KeyBackspace {
	//		p.backspace()
	//	}
	//})
	//p.Window.Resize(fyne.NewSize(500, 300))
	//p.Window.Show()
}
