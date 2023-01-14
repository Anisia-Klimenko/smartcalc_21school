package main

import (
	"calc/internal/about"
	"calc/internal/content"
	"calc/internal/history"
	"calc/internal/plot"
	"fyne.io/fyne/v2/app"
	"os"
	"strings"
)

func main() {
	// Get current location of executable to update assets path in case
	// app was installed as desktop one
	pwd, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}
	// App was installed
	if strings.Contains(pwd, ".app") {
		pwd = pwd[:strings.LastIndex(pwd, "/")]
		pwd += "/../../../../"
		about.UpdatePath(pwd)
		history.UpdatePaths(pwd)
		plot.UpdatePath(pwd)
	}

	// Init app
	a := app.New()
	// Create new calculator object
	c := content.NewCalculator(a)
	// Load UI and handlers to calculator object
	c.LoadUI(a)
	// Run app
	a.Run()
}
