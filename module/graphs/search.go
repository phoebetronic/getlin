package graphs

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	var out getlin.Vector
	{
		out = vec
	}

	for _, x := range m.mpr.All() {
		{
			m.cac.Create(out)
		}

		{
			out = vector.New(vector.Config{})
		}

		// Every layer x is comprised of modules which all receive the same
		// input of the preceeding layer. In case of the very first layer, this
		// input is the original input received during the call to
		// Module.Search. The search results of each layer x are concatenated so
		// that a consistent result of layer output vectors can be forwarded to
		// the next layer of modules.
		for _, y := range x {
			for _, b := range y.Search(m.cac.Latest().Inp).Bit() {
				out.Add(b)
			}
		}

		{
			m.cac.Update(out)
		}
	}

	// The search result of a layered architecture is always the last output
	// vector of the Graphs Module. Note that the output vector does not carry
	// the true labels of the original input vector, simply because these should
	// never be needed during inference. So we safe the complexity and compute.
	{
		out = m.cac.Latest().Out
	}

	{
		m.cac.Delete()
	}

	return out
}
