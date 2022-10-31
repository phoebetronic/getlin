package invert

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	var inp []bool
	{
		inp = vec.Inp().Raw()
	}

	for l, r := 0, len(inp)-1; l < r; l, r = l+1, r-1 {
		inp[l], inp[r] = inp[r], inp[l]
	}

	return vector.New(vector.Config{Out: inp})
}
