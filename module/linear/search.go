package linear

import "github.com/phoebetron/getlin/vector"

func (m *Module) Search(vec vector.Interface) vector.Interface {
	var prd vector.Interface
	{
		prd = vector.New(vector.Config{})
	}

	for _, x := range m.cla {
		prd.Add(x.Search(vec))
	}

	return prd
}
