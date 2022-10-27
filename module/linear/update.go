package linear

import "github.com/phoebetron/getlin/vector"

func (m *Module) Update(vec vector.Interface) {
	for i := range m.cla {
		var bck vector.Interface
		{
			bck = vector.New(vector.Config{
				Bit: vec.Bit(),
				Tru: []bool{vec.Tru(i)},
			})
		}

		{
			m.cla[i].Update(bck)
		}
	}
}
