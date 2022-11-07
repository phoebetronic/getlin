package multic

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	var out []uint8

	for _, x := range m.sin {
		out = append(out, x.Search(vec).Out().Raw()...)
	}

	return vector.New(vector.Config{Out: out})
}
