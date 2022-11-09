package single

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	if m.voting(vec.Inp()) >= 0 {
		return vector.New(vector.Config{Out: []float32{1}})
	} else {
		return vector.New(vector.Config{Out: []float32{0}})
	}
}

func (m *Module) voting(fea []float32) float32 {
	var out float32

	// All even Clauses have positive voting rights. Each of them may vote +1.
	for i := 0; i < len(m.cla); i += 2 {
		out += m.cla[i].Output(fea)
	}

	// All uneven Clauses have negative voting rights. Each of them may vote -1.
	for i := 1; i < len(m.cla); i += 2 {
		out -= m.cla[i].Output(fea)
	}

	return out
}
