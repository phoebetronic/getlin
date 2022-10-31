package metric

type states struct {
	// dat contains the distribution buckets of this state summary. In this
	// particular implementation we maintain 41 buckets for the full state
	// distribution. Buckets are split into 5% deltas on each side, plus one
	// bucket for 0.
	dat [41]int
}

func (s *states) Ind(rat float32) int {
	var hal int
	{
		hal = len(s.dat) / 2
	}

	var fac float32
	{
		fac = 1 / float32(hal)
	}

	if rat < 0 {
		return hal - int(rat/-fac)
	}

	if rat == 0 {
		return hal
	}

	// rat > 0
	return hal + int(rat/+fac)
}

func (s *states) Nrm() [41]float32 {
	var sum int

	for _, x := range s.dat {
		sum += x
	}

	if sum == 0 {
		return [41]float32{}
	}

	var nrm [41]float32
	for i := range s.dat {
		nrm[i] = float32(s.dat[i]) / float32(sum)
	}

	return nrm
}

func (s *states) Raw() [41]int {
	return s.dat
}
