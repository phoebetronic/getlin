package single

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	if m.voting(vec.Inp().Raw()) >= 0 {
		return vector.New(vector.Config{Out: []uint8{1}})
	} else {
		return vector.New(vector.Config{Out: []uint8{0}})
	}
}

func (m *Module) voting(fea []uint8) int {
	var out int

	// All even Clauses have positive voting rights. Each of them may vote +1.
	for i := 0; i < len(m.cla); i += 2 {
		out += int(m.cla[i].Output(fea))
	}

	// All uneven Clauses have negative voting rights. Each of them may vote -1.
	for i := 1; i < len(m.cla); i += 2 {
		out -= int(m.cla[i].Output(fea))
	}

	return out
}
