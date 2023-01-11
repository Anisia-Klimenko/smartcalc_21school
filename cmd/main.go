package main

import (
	"calc/internal/content"
	"fyne.io/fyne/v2/app"
)

func main() {
	//plot.TestPlot()
	a := app.New()
	c := content.NewCalculator(a)
	c.LoadUI(a)
	a.Run()
}
