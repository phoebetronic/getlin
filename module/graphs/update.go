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

	// The Vector Status is exclusively being initialized once at the very start
	// of the update process, regardless the depth and complexity of Module
	// architectures. This is to track the actual success or failure of system
	// performance during any given iteration. Success here means the predicted
	// output Vector matches the provided true labels.
	if !vec.Sta().Ini() {
		if vec.Out().Eql(m.cac.Lat().Out()) {
			vec.Sta().Upd(true)
		} else {
			vec.Sta().Upd(false)
		}
	}

	var tru []bool
	{
		tru = vec.Out().Raw()
	}

	for i, x := range m.mpr.All() {
		for j, y := range x {
			var out []bool
			if vec.Sta().Suc() {
				// In case of success, the true labels used for updating a
				// Module are those that the Module itself predicted during
				// inference earlier. This is the most specific kind of update
				// that can happen, and it is likely the most rare one too.
				{
					out = m.cac.Out(i)[j].Raw()
				}
			} else {
				// In case of failure, the true labels used for updating a
				// Module are those that the original input Vector carries. The
				// true label index range picks the correct partial of the true
				// labels for the updating process.
				var ind [2]int
				{
					ind = m.mpr.Tru(y)
				}

				var prt []bool
				{
					prt = tru[ind[0]:ind[1]]
				}

				// In case the true labels are picked in the scope of a Voting
				// Module the single true bit has to be scaled to the output
				// Vector length of the Module to be updated. For any other
				// partial shapes the content negotiation assumes y to be a
				// Linear Module which can receive the full true label partial
				// as is.
				if len(prt) == 1 {
					out = vector.Repeat(prt, y.Shaper().Out())
				} else {
					out = prt
				}
			}

			{
				y.Update(vector.New(vector.Config{
					Inp: m.cac.Vec(i).Inp().Raw(),
					Out: out,
					Sta: vec.Sta(),
				}))
			}
		}
	}

	{
		m.cac.Del()
	}
}
