package invert

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	var bit []bool
	{
		bit = vec.Bit()
	}

	for l, r := 0, len(bit)-1; l < r; l, r = l+1, r-1 {
		bit[l], bit[r] = bit[r], bit[l]
	}

	var tru []bool
	{
		tru = vec.Tru()
	}

	for l, r := 0, len(tru)-1; l < r; l, r = l+1, r-1 {
		tru[l], tru[r] = tru[r], tru[l]
	}

	var inv getlin.Vector
	{
		inv = vector.New(vector.Config{
			Bit: bit,
			Tru: tru,
		})
	}
	return inv
}
