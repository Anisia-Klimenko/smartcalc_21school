package main

import (
	"calc/internal/about"
	"calc/internal/content"
	"calc/internal/history"
	"fmt"
	"fyne.io/fyne/v2/app"
	"os"
	"strings"
)

func main() {
	pwd, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.Contains(pwd, ".app") {
		pwd = pwd[:strings.LastIndex(pwd, "/")]
		pwd += "/../../../../"
		about.UpdatePath(pwd)
		history.UpdatePaths(pwd)
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
