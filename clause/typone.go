package clause

func (c *Clause) TypOne(fea []float32, prd float32) {
	for i := range fea {
		// Type 1a Feedback reinforces includes of the feature weights so that
		// the Clause itself gradually resembles the feature weights of the
		// pattern it recognized.
		if prd > 0 {
			if fea[i] > 0.5 {
				// A Clause prediction above 0 means some pattern was
				// recognized. To further increase the likelyhood of inclusion
				// the negative polarity feature weights above 0.5 are decreased
				// with low probability.
				if !c.neg[i].Low() && c.actlow() {
					c.neg[i].Red(c.stpsiz())
				}
				// A Clause prediction above 0 means some pattern was
				// recognized. To further increase the likelyhood of inclusion
				// the positive polarity feature weights above 0.5 are increased
				// with high probability.
				if !c.pos[i].Hig() && c.acthig() {
					c.pos[i].Add(c.stpsiz())
				}
			} else {
				// A Clause prediction above 0 means some pattern was
				// recognized. To further increase the likelyhood of inclusion
				// the negative polarity feature weights below 0.5 are increased
				// with high probability.
				if !c.neg[i].Hig() && c.acthig() {
					c.neg[i].Add(c.stpsiz())
				}
				// A Clause prediction above 0 means some pattern was
				// recognized. To further increase the likelyhood of inclusion
				// the positive polarity feature weights below 0.5 are decreased
				// with low probability.
				if !c.pos[i].Low() && c.actlow() {
					c.pos[i].Red(c.stpsiz())
				}
			}
		}
		// Type 1b Feedback combats over-fitting by reinforcing excludes of the
		// feature weights in case the Clause did not recognize any pattern.
		if prd == 0 {
			// A Clause prediction of 0 means no pattern was recognized. To
			// further reduce the likelyhood of inclusion all negative polarity
			// states are getting decreased stochastically with low probability,
			if !c.neg[i].Low() && c.actlow() {
				c.neg[i].Red(c.stpsiz())
			}
			// A Clause prediction of 0 means no pattern was recognized. To
			// further reduce the likelyhood of inclusion all positive polarity
			// states are getting decreased stochastically with low probability,
			if !c.pos[i].Low() && c.actlow() {
				c.pos[i].Red(c.stpsiz())
			}
		}
	}
}

func (c *Clause) actlow() bool {
	return c.ran.F32() <= c.fre
}

func (c *Clause) acthig() bool {
	return c.ran.F32() <= 1-c.fre
}

func (c *Clause) stpsiz() float32 {
	return 0.1
}
