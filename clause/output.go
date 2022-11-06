package clause

func (c *Clause) Output(fea []uint8) uint8 {
	for i := range fea {
		// If any NEGATIVE polarity 1 is included we return 0 immediately, since
		// the required ANDing conjunction can never be true.
		if fea[i] == 1 && c.neg[i].Inc() {
			return 0
		}

		// If any POSITIVE polarity 0 is included we return 0 immediately, since
		// the required ANDing conjunction can never be true.
		if fea[i] == 0 && c.pos[i].Inc() {
			return 0
		}
	}

	// The Clause successfully returns 1 if neither negative 1 nor positive 0 is
	// included, given the current feature Vector.
	return 1
}
