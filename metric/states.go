package metric

type states struct {
	// dat contains the distribution buckets of this state summary. In this
	// particular implementation we maintain 41 buckets for the full state
	// distribution. Buckets are split into 5% deltas on each side, plus one
	// bucket for 0.
	dat [41]float32
}

func (s *states) Nrm() [41]float32 {
	var sum float32

	for _, x := range s.dat {
		sum += x
	}

	var nrm [41]float32
	for i := range s.dat {
		nrm[i] = s.dat[i] / sum
	}

	return nrm
}

func (s *states) Raw() [41]float32 {
	return s.dat
}
