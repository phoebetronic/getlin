package metric

const (
	TP = 0
	TN = 1
	FP = 2
	FN = 3
)

type matrix struct {
	// dat contains the TP, TN, FP and FN counts at index 0, 1, 2 and 3
	// respectively.
	dat [4]float32
}

func (m *matrix) Acc() float32 {
	return (m.dat[TP] + m.dat[TN]) / (m.dat[TP] + m.dat[TN] + m.dat[FP] + m.dat[FN])
}

func (m *matrix) Mae() float32 {
	return (m.dat[FP] + m.dat[FN]) / (m.dat[TP] + m.dat[TN] + m.dat[FP] + m.dat[FN])
}

func (m *matrix) Ppv() float32 {
	return m.dat[TP] / (m.dat[TP] + m.dat[FP])
}
