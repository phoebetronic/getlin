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
			var t []int
			{
				t = m.mpr.Tru(y)
			}

			//  Purely functional modules cannot be updated since they have no
			//  learnable component within them. For those modules the Mapper
			//  does not provide a true label index range. So we skip them to
			//  ensure only learnable modules are being updated.
			if len(t) == 0 {
				continue
			}

			{
				y.Update(vector.New(vector.Config{
					Inp: m.cac.Vec(i).Inp().Raw(),
					Out: tru[t[0]:t[1]],
				}))
			}
		}
	}

	{
		m.cac.Del()
	}
}
