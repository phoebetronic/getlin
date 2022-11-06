package clause

func (c *Clause) TypTwo(fea []uint8, prd uint8) {
	if prd == 1 {
		for i := range fea {
			// In case the Clause's label prediction is 1 and the negative
			// polarity input feature 1 is excluded, the negative polarity
			// states are getting increased.
			if fea[i] == 1 && c.neg[i].Exc() {
				c.neg[i].Add()
			}
			// In case the Clause's label prediction is 1 and the positive
			// polarity input feature 0 is excluded, the positive polarity
			// states are getting increased.
			if fea[i] == 0 && c.pos[i].Exc() {
				c.pos[i].Add()
			}
		}
	}
}
