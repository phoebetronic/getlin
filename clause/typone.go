package clause

func (c *Clause) TypOne(fea []float32, prd float32) {
	if prd == 0 {
		for i := range fea {
			// A Clause prediction of 0 means no pattern was recognized. To
			// further reduce the likelyhood of inclusion all negative polarity
			// states are getting decreased stochastically with low probability,
			if !c.neg[i].Low() && c.actlow() {
				c.neg[i].Red(0.1)
			}
			// A Clause prediction of 0 means no pattern was recognized. To
			// further reduce the likelyhood of inclusion all positive polarity
			// states are getting decreased stochastically with low probability,
			if !c.pos[i].Low() && c.actlow() {
				c.pos[i].Red(0.1)
			}
		}
	} else {
		for i := range fea {
			if fea[i] <= 0.5 {
				// In case the Clause's label prediction is 1 and the input
				// feature is 0, the negative polarity states are getting
				// increased stochastically with high probability,
				if !c.neg[i].Hig() && c.acthig() {
					c.neg[i].Add(0.1)
				}
				// In case the Clause's label prediction is 1 and the input
				// feature is 0, the positive polarity states are getting
				// decreased stochastically with low probability,
				if !c.pos[i].Low() && c.actlow() {
					c.pos[i].Red(0.1)
				}
			} else {
				// In case the Clause's label prediction is 1 and the input
				// feature is 1, the negative polarity states are getting
				// decreased stochastically with low probability,
				if !c.neg[i].Low() && c.actlow() {
					c.neg[i].Red(0.1)
				}
				// In case the Clause's label prediction is 1 and the input
				// feature is 1, the positive polarity states are getting
				// increased stochastically with high probability,
				if !c.pos[i].Hig() && c.acthig() {
					c.pos[i].Add(0.1)
				}
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
