package voting

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Update(vec getlin.Vector) {
	if m.wei == nil {
		return
	}

	// The weights Module is conditioned based on the input patterns to find a
	// threshold for. The true label received is being scaled to the output
	// Vector lenght defined. Patterns supposed to match will increase the
	// majority threshold with every round of feedback, while patterns supposed
	// not to match will decrease that threshold.
	{
		m.wei.Update(vector.New(vector.Config{
			Inp: vec.Inp().Raw(),
			Out: vector.Repeat(vec.Out().Raw(), m.wei.Shaper().Out()),
			Sta: vec.Sta(),
		}))
	}
}
