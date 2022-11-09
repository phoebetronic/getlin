package clause

func (c *Clause) TypTwo(fea []float32, prd float32) {
	if prd > 0 {
		for i := range fea {
			// A Clause prediction above 0 means some pattern was recognized. To
			// further increase the likelyhood of inclusion the now excluded
			// negative polarity input features above 0.5 are getting increased.
			if fea[i] > 0.5 && c.neg[i].Exc() {
				c.neg[i].Add(0.1)
			}
			// A Clause prediction above 0 means some pattern was recognized. To
			// further increase the likelyhood of inclusion the now excluded
			// positive polarity input features below 0.5 are getting increased.
			if fea[i] <= 0.5 && c.pos[i].Exc() {
				c.pos[i].Add(0.1)
			}
		}
	}
}
