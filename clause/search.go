package clause

import "github.com/phoebetron/getlin/vector"

func (c *Clause) Search(vec vector.Interface) bool {
	for i := 0; i < vec.Len(); i++ {
		if c.neg[i].Inc() && !vec.Neg(i) {
			return false
		}
		if c.pos[i].Inc() && !vec.Pos(i) {
			return false
		}
	}

	return true
}
