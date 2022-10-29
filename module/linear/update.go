package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Update(vec getlin.Vector) {
	for i := range m.cla {
		var bck getlin.Vector
		{
			bck = vector.New(vector.Config{
				Bit: vec.Bit(),
				Tru: []bool{vec.Tru()[i]},
			})
		}

		{
			m.cla[i].Update(bck)
		}
	}
}
