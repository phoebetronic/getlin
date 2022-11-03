package graphs

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	var out getlin.Vector
	{
		out = m.search(vec)
	}

	{
		m.cac.Del()
	}

	return out
}

func (m *Module) search(vec getlin.Vector) getlin.Vector {
	for _, x := range m.mpr.All() {
		// The first call to Cacher.Add in every round records a new linker for
		// the current layer. Since we use the output vector of the previous
		// iteration as input for the current layer, we need to register a copy
		// so that further usage of the Cacher during e.g. Module.Update does
		// not lead to any unintended side effects.
		{
			m.cac.Add(vec.Inp())
		}

		// Every layer x is comprised of modules which all receive the same
		// input of the preceeding layer. In case of the very first layer, this
		// input is the original input Vector received during the call to
		// Module.Search. The search results of each layer x are concatenated so
		// that a consistent result of layer output vectors can be forwarded to
		// the next layer of modules. Below the Cacher collects the Module
		// search results for the whole layer currently being processed.
		for _, y := range x {
			m.cac.Upd(y.Search(m.cac.Lat()).Out())
		}

		// The latest output Vector is now the complete concatenated bit array
		// produced by each and every Module of layer x. That output becomes the
		// input for the next iteration feeding into all of the modules within
		// the following layer.
		{
			vec = vector.New(vector.Config{Inp: m.cac.Lat().Out().Raw()})
		}
	}

	// The search result of a layered architecture is always the last output
	// vector of the Graphs Module. Note that the output vector does not carry
	// the true labels of the original input vector, simply because these should
	// never be needed during inference. So we safe the complexity and compute.
	return m.cac.Lat()
}
