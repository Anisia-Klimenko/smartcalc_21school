package content

import (
	"calc/internal/math"
	"log"
)

func (c *Calc) clear() {
	c.display("")
}

func (c *Calc) evaluate() {
	log.Println(c.Output.Text)
	c.display(math.Calculate(c.Output.Text))
	//if strings.Contains(c.Output.Text, "error") {
	//	c.display("error")
	//	return
	//}
	//
	//expression, err := govaluate.NewEvaluableExpression(c.Output.Text)
	//if err != nil {
	//	log.Println("Error in calculation", err)
	//	c.display("error")
	//	return
	//}
	//
	//result, err := expression.Evaluate(nil)
	//if err != nil {
	//	log.Println("Error in calculation", err)
	//	c.display("error")
	//	return
	//}
	//
	//value, ok := result.(float64)
	//if !ok {
	//	log.Println("Invalid input:", c.Output.Text)
	//	c.display("error")
	//	return
	//}
	//
	//c.display(strconv.FormatFloat(value, 'f', -1, 64))
}
