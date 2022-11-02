package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Update(vec getlin.Vector) {
	for i := range m.cla {
		m.cla[i].Update(vector.New(vector.Config{
			Inp: vec.Inp().Raw(),
			Out: []bool{vec.Out().Raw()[i]},
		}))
	}
}
