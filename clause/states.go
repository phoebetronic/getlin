package clause

func (c *Clause) States() ([]float32, []float32) {
	var neg []float32
	for _, x := range c.neg {
		neg = append(neg, x.Poi())
	}

	var pos []float32
	for _, x := range c.pos {
		pos = append(pos, x.Poi())
	}

	return neg, pos
}
