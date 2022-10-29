package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	var prd getlin.Vector
	{
		prd = vector.New(vector.Config{})
	}

	for _, x := range m.cla {
		prd.Add(x.Search(vec))
	}

	return prd
}
