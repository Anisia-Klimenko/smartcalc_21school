package content

import (
	"calc/internal/math"
	"log"
)

func (c *Calc) clear() {
	c.display("")
}

func (c *Calc) mod() {
	c.display("")
}

func (c *Calc) evaluate() {
	log.Println(c.Output.Text)
	c.display(math.Calculate(c.Output.Text, c.XValue))
}
