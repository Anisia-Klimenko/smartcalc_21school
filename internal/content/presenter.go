package content

import (
	"calc/internal/history"
	"calc/internal/math"
	"log"
)

func (c *Calc) clear() {
	c.display("")
}

func (c *Calc) evaluate() {
	var historyText = c.Output.Text + "="
	c.ifEqualPressed = true
	log.Println(c.Output.Text)
	result := math.Calculate(c.Output.Text, c.XValue)
	historyText += result
	history.SaveHistory(historyText)
	c.display(result)
}
