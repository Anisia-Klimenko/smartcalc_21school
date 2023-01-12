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
	result := math.Calculate(c.Output.Text, c.XValue)
	historyText += result
	history.UpdateHistory(historyText)
	c.display(result)
	log.Println(historyText)
}
