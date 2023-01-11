package plot

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/ajstarks/fc"
	"github.com/ajstarks/fc/chart"
	"image/color"
	"io"
	"os"
)

func TestPlot(a fyne.App) {
	var input io.Reader
	var ferr error

	// Read from specified file
	input, ferr = os.Open("../assets/data.d")
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", ferr)
		os.Exit(1)
	}

	// Read in the data
	plot, err := chart.DataRead(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	// Define the field
	w, h := 600, 600
	r := &canvas.Rectangle{FillColor: color.RGBA{A: 255}}
	r.Move(fyne.Position{X: float32(w/2) - (float32(w) / 2), Y: float32(h/2) - (float32(h) / 2)})
	r.Resize(fyne.Size{Width: float32(w), Height: float32(h)})
	field := fc.Canvas{
		Window:    a.NewWindow(fmt.Sprintf("Chart: %s", plot.Title)),
		Container: fyne.NewContainerWithoutLayout(r),
		Width:     float64(w),
		Height:    float64(h),
	}
	// Define the colors
	datacolor := fc.ColorLookup("steelblue")
	labelcolor := color.RGBA{100, 100, 100, 255}
	bgcolor := color.RGBA{255, 255, 255, 255}
	field.Rect(50, 50, 100, 100, bgcolor)

	// Set the plot attributes
	plot.Zerobased = true

	// Draw the data
	plot.Color = datacolor
	plot.Frame(field, 1)
	plot.Line(field, 0.25)
	plot.Bar(field, 0.25)
	plot.Scatter(field, 0.5)

	// Draw labels, axes if specified
	plot.Color = labelcolor
	plot.Label(field, 1.5, 5)
	var yaxmin, yaxmax, yaxstep float64
	if n, err := fmt.Sscanf("-1,1,0.25", "%v,%v,%v", &yaxmin, &yaxmax, &yaxstep); n == 3 && err == nil {
		plot.YAxis(field, 1.5, yaxmin, yaxmax, yaxstep, "%v", true)
	}

	field.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			field.Window.Close()
		}
	})

	// Show the plot
	field.Window.Resize(fyne.NewSize(600, 600))
	field.Window.SetFixedSize(true)
	field.Window.SetPadded(false)
	field.Window.SetContent(field.Container)
	field.Window.Show()
}
