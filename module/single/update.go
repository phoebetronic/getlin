package single

import (
	"github.com/phoebetron/getlin"
)

func (m *Module) Update(vec getlin.Vector) {
	var fea []float32
	{
		fea = vec.Inp()
	}

	var tru float32
	{
		tru = vec.Out()[0]
	}

	var prd []float32
	{
		prd = make([]float32, len(m.cla))
	}

	var vot float32

	for i := 0; i < len(m.cla); i += 2 {
		var out float32
		{
			out = m.cla[i].Output(fea)
		}

		{
			prd[i] = out
		}

		{
			vot += out
		}
	}

	for i := 1; i < len(m.cla); i += 2 {
		var out float32
		{
			out = m.cla[i].Output(fea)
		}

		{
			prd[i] = out
		}

		{
			vot -= out
		}
	}

	if tru <= 0.5 {
		for i := 0; i < len(m.cla); i += 2 {
			if m.actzer(vot) {
				m.cla[i].TypTwo(fea, prd[i])
			}
		}
		for i := 1; i < len(m.cla); i += 2 {
			if m.actzer(vot) {
				m.cla[i].TypOne(fea, prd[i])
			}
		}
	} else {
		for i := 0; i < len(m.cla); i += 2 {
			if m.actone(vot) {
				m.cla[i].TypOne(fea, prd[i])
			}
		}
		for i := 1; i < len(m.cla); i += 2 {
			if m.actone(vot) {
				m.cla[i].TypTwo(fea, prd[i])
			}
		}
	}
}

func (m *Module) actone(vot float32) bool {
	// In case the negative voting threshold is reached for the positive
	// prediction, the Clause's TAs are more or less useless and feedback is
	// fully enabled without having to access expensive entropy devices.
	if vot <= -1 {
		return true
	}

	// In case the positive voting threshold is reached for the positive
	// prediction, the Clause's TAs are sufficiently synced and feedback is
	// completely disabled without having to access expensive entropy devices.
	if vot >= +1 {
		return false
	}

	// Within the voting threshold range, more feedback is given the more
	// negative the prediction is.
	return m.ran.F32() <= (1-vot)/(2)
}

func (m *Module) actzer(vot float32) bool {
	// In case the negative voting threshold is reached for the negative
	// prediction, the Clause's TAs are sufficiently synced and feedback is
	// completely disabled without having to access expensive entropy devices.
	if vot <= -1 {
		return false
	}

	// In case the positive voting threshold is reached for the negative
	// prediction, the Clause's TAs are more or less useless and feedback is
	// fully enabled without having to access expensive entropy devices.
	if vot >= +1 {
		return true
	}

	// Within the voting threshold range, more feedback is given the more
	// positive the prediction is.
	return m.ran.F32() <= (1+vot)/(2)
}
