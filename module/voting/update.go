package voting

import (
	"github.com/phoebetron/getlin"
)

func (m *Module) Update(vec getlin.Vector) {
	var fea []uint8
	{
		fea = vec.Inp().Raw()
	}

	var tru uint8
	{
		tru = vec.Out().Raw()[0]
	}

	var out []uint8
	{
		out = make([]uint8, len(m.cla))
	}

	var vot float32

	// All even Clauses have positive voting rights. Each of them may vote +1.
	for i := 0; i < len(m.cla); i += 2 {
		var prd uint8
		{
			prd = m.cla[i].Output(fea)
		}

		{
			out[i] = prd
		}

		{
			vot += float32(prd)
		}
	}

	// All uneven Clauses have negative voting rights. Each of them may vote -1.
	for i := 1; i < len(m.cla); i += 2 {
		var prd uint8
		{
			prd = m.cla[i].Output(fea)
		}

		{
			out[i] = prd
		}

		{
			vot -= float32(prd)
		}
	}

	if tru == 0 {
		for i := 0; i < len(m.cla); i += 2 {
			if m.actcla(vot) {
				m.cla[i].TypTwo(fea, out[i])
			}
		}
		for i := 1; i < len(m.cla); i += 2 {
			if m.actcla(vot) {
				m.cla[i].TypOne(fea, out[i])
			}
		}
	} else {
		for i := 0; i < len(m.cla); i += 2 {
			if m.actcla(vot) {
				m.cla[i].TypOne(fea, out[i])
			}
		}
		for i := 1; i < len(m.cla); i += 2 {
			if m.actcla(vot) {
				m.cla[i].TypTwo(fea, out[i])
			}
		}
	}
}

func (m *Module) actcla(vot float32) bool {
	return m.ran.F32() <= (m.thr+vot)/(2*m.thr)
}
