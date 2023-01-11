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
	c.Window = c.App.NewWindow("Plot (c) acristin")
}

func (p *Plot) LoadPlotWindow(a fyne.App) {
}
