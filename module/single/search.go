package single

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	return vector.New(vector.Config{Out: []float32{m.output(vec.Inp())}})
}

func (m *Module) output(fea []float32) float32 {
	var out float32

	// All even Clauses have positive voting rights. Each of them may vote +1.
	for i := 0; i < len(m.cla); i += 2 {
		out += m.cla[i].Output(fea)
	}

	// All uneven Clauses have negative voting rights. Each of them may vote -1.
	for i := 1; i < len(m.cla); i += 2 {
		out -= m.cla[i].Output(fea)
	}

	// A ReLu clipping is applied to the returned voting result so that the
	// output Vector range is within [0.0, 1.0].
	//
	//                   •
	//                   1         + + +
	//                   •       +
	//                   •     +
	//                   •   +
	//                   • +
	//     + + 1 + + + + 0 • • • • 1 • •
	//                   •
	//                   •
	//                   •
	//                   •
	//                   1
	//                   •
	//
	return minf32(maxf32(0, out), 1)
}

func maxf32(a float32, b float32) float32 {
	if a > b {
		return a
	}

	return b
}

func minf32(a float32, b float32) float32 {
	if a < b {
		return a
	}

	return b
}
