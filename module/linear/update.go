package linear

import (
	"github.com/phoebetron/getlin"
)

func (m *Module) Update(vec getlin.Vector) {
	for i := range m.cla {
		m.cla[i].Update(vec)
	}
}
