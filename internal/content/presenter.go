package content

import (
	"calc/internal/math"
	"log"
)

func (c *Calc) clear() {
	c.display("")
}

func (c *Calc) evaluate() {
	c.ifEqualPressed = true
	log.Println(c.Output.Text)
	c.display(math.Calculate(c.Output.Text, c.XValue))
}
