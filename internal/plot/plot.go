package plot

import (
	"calc/internal/file"
	"calc/internal/math"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/ajstarks/fc"
	"github.com/ajstarks/fc/chart"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
)

type Borders struct {
	XMin float64
	XMax float64
	YMin float64
	YMax float64
}

// UpdateData generates data file for plotting
func UpdateData(equation string, xMin, xMax, yMin, yMax float64) (float64, float64) {
	var newMin, newMax = yMax, yMin
	dataFile := "../assets/data.d"

	// Clear file before writing
	file.Clear(dataFile)

	// Set plot title
	file.Update(dataFile, "# y="+equation)

	// Generate dataset
	delta := (xMax - xMin) / 100
	for i := xMin; i < xMax; i += delta {
		result, err := strconv.ParseFloat(math.Calculate(equation, i), 64)
		if err == nil {
			// Handle borders
			if result > yMax {
				result = yMax
			} else if result < yMin {
				result = yMin
			}
			// Find new yMax and yMin
			if result > newMax {
				newMax = result
			}
			if result < newMin {
				newMin = result
			}
			// Check if result is not NaN
			if result == result {
				file.Update(dataFile, fmt.Sprintf("%.2f\t%.7f", i, result))
			}
		}
	}
	return newMin, newMax
}

func ShowPlot(a fyne.App, equation string, border Borders) {
	var input io.Reader
	var ferr error
	// Handle empty equation
	if len(equation) == 0 {
		log.Println("plot: empty equation")
		return
	}
	log.Println("plot: opened")

	// Generate dataset and save to file, update y-borders
	border.YMin, border.YMax = UpdateData(equation, border.XMin, border.XMax, border.YMin, border.YMax)

	// Read from specified file
	input, ferr = os.Open("../assets/data.d")
	if ferr != nil {
		log.Println("plot: data file open error")
		return
	}

	// Read dataset
	plot, err := chart.DataRead(input)
	if err != nil {
		log.Println("plot: data file read error")
		return
	}
	log.Println("plot:", plot.Title)

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
	dataColor := fc.ColorLookup("steelblue")
	labelColor := color.RGBA{R: 100, G: 100, B: 100, A: 255}
	bgColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	field.Rect(50, 50, 100, 100, bgColor)

	// Set the plot attributes
	plot.Zerobased = true
	plot.Color = dataColor
	plot.Line(field, 0.25)
	plot.Bar(field, 0.25)
	plot.Scatter(field, 0.25)

	// Draw labels and axes
	plot.Color = labelColor
	plot.Label(field, 1.5, 5)
	plot.YAxis(field, 1.5, border.YMin, border.YMax, (border.YMax-border.YMin)/20, "%.2f", true)

	// Handle shortcuts
	field.Window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape || keyEvent.Name == "W" {
			// Close window in case ESC or W was pressed
			field.Window.Close()
			log.Println("plot: closed")
		}
	})

	// Show the plot window
	field.Window.Resize(fyne.NewSize(600, 600))
	//field.Window.SetPadded(true)
	field.Window.SetContent(field.Container)
	field.Window.Show()
}
