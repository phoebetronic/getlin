package clause

func (c *Clause) States() ([]int, []int) {
	var neg []int
	for _, x := range c.neg {
		neg = append(neg, x.Poi())
	}

	var pos []int
	for _, x := range c.pos {
		pos = append(pos, x.Poi())
	}

	return neg, pos
}
