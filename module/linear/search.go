package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	var out getlin.Vector
	{
		out = vector.New(vector.Config{})
	}

	for _, x := range m.cla {
		out.Out().Add(x.Search(vec))
	}

	return out
}
