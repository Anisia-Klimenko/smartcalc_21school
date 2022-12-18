package main

import (
	"calc/internal/content"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	c := content.NewCalculator()
	c.LoadUI(a)
	a.Run()
}
