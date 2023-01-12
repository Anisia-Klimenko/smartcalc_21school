package content

import (
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// changeXValue convert entry input s to float64 and saves it in val
func (c *Calc) changeXValue(s string, val *float64) {
	*val, _ = strconv.ParseFloat(s, 64)
}

// addEntry creates a new single line entry widget with the set label
// and onchange handler.
func (c *Calc) addEntry(label string, handler func(string)) *widget.Entry {
	entry := widget.NewEntry()
	entry.TextStyle.Monospace = true
	entry.OnChanged = handler
	entry.PlaceHolder = label
	return entry
}
