package clause

import "github.com/phoebetron/getlin"

func (c *Clause) Search(vec getlin.Vector) bool {
	for i := range vec.Bit() {
		if c.neg[i].Inc() && !vec.Neg(i) {
			return false
		}
		if c.pos[i].Inc() && !vec.Pos(i) {
			return false
		}
	}

	return true
}
