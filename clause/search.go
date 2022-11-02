package clause

import "github.com/phoebetron/getlin"

func (c *Clause) Search(vec getlin.Vector) bool {
	var one int

	for i := range vec.Inp().Raw() {
		if c.neg[i].Inc() {
			if vec.Inp().Neg(i) {
				one++
			} else {
				// If any 0 is included in the negative polarity array we return
				// 0 immediately, since the required ANDing conjunction can
				// never be true.
				return false
			}
		}
		if c.pos[i].Inc() {
			if vec.Inp().Pos(i) {
				one++
			} else {
				// If any 0 is included in the positive polarity array we return
				// 0 immediately, since the required ANDing conjunction can
				// never be true.
				return false
			}
		}
	}

	// If we have at least a single true bit included, while not having any
	// false bit included, we return 1.
	return one > 0
}
