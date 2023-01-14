package content

import "log"

// isInLimits checks if val is in [-1000000; 1000000]
func isInLimits(val float64) bool {
	return val >= -1000000 && val <= 1000000
}

// checkBorders checks if borders are valid and if not changes them to default
// valid values
func (c *Calc) checkBorders() (ifValid bool) {
	// x min border is not set
	if !c.IsBorderSet.XMin {
		c.Border.XMin = -10
	}
	// x max border is not set
	if !c.IsBorderSet.XMax {
		c.Border.XMax = 10
	}
	// y min border is not set
	if !c.IsBorderSet.YMin {
		c.Border.YMin = -30
	}
	// y max border is not set
	if !c.IsBorderSet.YMax {
		c.Border.YMax = 30
	}

	ifValid = true
	log.Println("plot: borders: [", c.Border.XMin, ";", c.Border.XMax, "] [", c.Border.YMin, ";", c.Border.YMax, "]")

	// non-valid borders
	if c.Border.XMin >= c.Border.XMax || c.Border.YMin >= c.Border.YMax ||
		!isInLimits(c.Border.XMin) || !isInLimits(c.Border.XMax) ||
		!isInLimits(c.Border.YMin) || !isInLimits(c.Border.YMax) {
		log.Println("plot: non-valid borders")
		c.Border.XMin = 0
		c.Border.YMin = 0
		c.Border.XMax = 0
		c.Border.YMax = 0
		c.display("non-valid borders")
		ifValid = false
	}
	return
}
