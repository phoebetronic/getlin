package clause

import "github.com/phoebetron/getlin"

func (c *Clause) Search(vec getlin.Vector) bool {
	for i := range vec.Inp().Raw() {
		if c.neg[i].Inc() && !vec.Inp().Neg(i) {
			return false
		}
		if c.pos[i].Inc() && !vec.Inp().Pos(i) {
			return false
		}
	}

	return true
}
