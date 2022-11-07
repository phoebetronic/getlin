package multic

import (
	"github.com/phoebetron/getlin"
)

func (m *Module) Update(vec getlin.Vector) {
	m.sin[vec.Cla()].Update(vec)
}
