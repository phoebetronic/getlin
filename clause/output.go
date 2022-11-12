package clause

func (c *Clause) Output(fea []float32) float32 {
	var sum float32

	for i := range fea {
		if c.neg[i].Inc() {
			if fea[i] > 0.5 {
				return 0
			} else {
				sum += c.neg[i].Poi()
			}
		}

		if c.pos[i].Inc() {
			if fea[i] <= 0.5 {
				return 0
			} else {
				sum += c.pos[i].Poi()
			}
		}
	}

	return sum / float32(len(fea)*2)
}
