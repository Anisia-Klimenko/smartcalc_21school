package main

import (
	"calc/internal/content"
	"fyne.io/fyne/v2/app"
)

func main() {
	// Init app
	a := app.New()
	// Create new calculator object
	c := content.NewCalculator(a)
	// Load UI and handlers to calculator object
	c.LoadUI(a)
	// Run app
	a.Run()
}
