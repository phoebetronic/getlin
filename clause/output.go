package clause

func (c *Clause) Output(fea []float32) float32 {
	var sum float32

	for i := range fea {
		// If any NEGATIVE polarity 1 is included we return 0 immediately, since
		// the required ANDing conjunction can never be true.
		if c.neg[i].Inc() {
			if fea[i] > 0.5 {
				return 0
			} else {
				sum += c.neg[i].Poi()
			}
		}

		// If any POSITIVE polarity 0 is included we return 0 immediately, since
		// the required ANDing conjunction can never be true.
		if c.pos[i].Inc() {
			if fea[i] <= 0.5 {
				return 0
			} else {
				sum += c.pos[i].Poi()
			}
		}
	}

	// The Clause successfully returns 1 if neither negative 1 nor positive 0 is
	// included, given the current feature Vector.
	return sum / float32(len(fea)*2)
}
