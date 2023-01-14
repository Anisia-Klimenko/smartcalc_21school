package content

import (
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// changeXValue is x-value handler, it converts entry input str to float64
// and saves it in val
func (c *Calc) changeXValue(str string, val *float64) {
	*val, _ = strconv.ParseFloat(str, 64)
}

// changeBorderValue is plot borders handler, it converts entry input str to
// float64 and saves it in val
func (c *Calc) changeBorderValue(str string, val *float64, isSet *bool) {
	// Set flag isSet only in case of non-empty input. This flag helps distinguish
	// between the cases where the input is 0 and when the input is empty.
	if len(str) > 0 {
		*isSet = true
	}
	*val, _ = strconv.ParseFloat(str, 64)
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
