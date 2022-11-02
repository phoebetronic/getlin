package graphs

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Update(vec getlin.Vector) {
	// Inference is required for updating the Graphs Module. The internal search
	// goes through all modules of all layers, filling the internal vector cache
	// so that we know on which level we produce which input and which output
	// given the current state of the system.
	{
		m.search(vec)
	}

	var tru []bool
	{
		tru = vec.Out().Raw()
	}

	for i, x := range m.mpr.All() {
		for _, y := range x {
			// The true label range indices are used to create the desired
			// output Vector for the next Module being updated. The output
			// Vector we generate might be the scaled version of a single bit.
			// That is, in nested architectures a whole Graphs Module may
			// receive a single true label, which then has to be applied during
			// feedback in any Module being part of its parent. The nested
			// modules then may define multiple outputs themselves, which means
			// that, the single true label received via its parent must be
			// scaled to meet every nested Module's output requirements, in
			// order for them to be updated.
			var t [2]int
			{
				t = m.mpr.Tru(y)
			}

			{
				y.Update(vector.New(vector.Config{
					Inp: m.cac.Vec(i).Inp().Raw(),
					Out: vector.Repeat(tru[t[0]:t[1]], y.Shaper().Out()),
				}))
			}
		}
	}

	{
		m.cac.Del()
	}
}
