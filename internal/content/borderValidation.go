package content

// isInLimits checks if val is in [-1000000; 1000000]
func isInLimits(val float64) bool {
	if val >= -1000000 && val <= 1000000 {
		return true
	}
	return false
}

// checkBorders checks if borders are valid and if not changes them to default
// valid values
func (c *Calc) checkBorders() {
	// x borders are not set
	if c.Border.XMin == 0 && c.Border.XMax == 0 {
		c.Border.XMin = -10
		c.Border.XMax = 10
	}
	// y borders are not set
	if c.Border.YMin == 0 && c.Border.YMax == 0 {
		c.Border.YMin = -30
		c.Border.YMax = 30
	}
	// non-valid borders
	if c.Border.XMin >= c.Border.XMax || c.Border.YMin >= c.Border.YMax ||
		!isInLimits(c.Border.XMin) || !isInLimits(c.Border.XMax) ||
		!isInLimits(c.Border.YMin) || !isInLimits(c.Border.YMax) {
		c.Border.XMin = 0
		c.Border.YMin = 0
		c.Border.XMax = 0
		c.Border.YMax = 0
		c.display("non valid borders")
	}
}
