package clause

func (c *Clause) States() []float32 {
	var sta []float32

	for _, x := range c.neg {
		sta = append(sta, x.Poi())
	}

	for _, x := range c.pos {
		sta = append(sta, x.Poi())
	}

	return sta
}
