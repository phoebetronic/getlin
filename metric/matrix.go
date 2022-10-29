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
	dat [4]int
}

func (m *matrix) Acc() float32 {
	return float32(m.dat[TP]+m.dat[TN]) / float32(m.dat[TP]+m.dat[TN]+m.dat[FP]+m.dat[FN])
}

func (m *matrix) Mae() float32 {
	return float32(m.dat[FP]+m.dat[FN]) / float32(m.dat[TP]+m.dat[TN]+m.dat[FP]+m.dat[FN])
}

func (m *matrix) Ppv() float32 {
	return float32(m.dat[TP]) / float32(m.dat[TP]+m.dat[FP])
}

func (m *matrix) Raw() [4]int {
	return m.dat
}
