package metric

type setter struct {
	sha *shared
}

func (s *setter) Mat(ind int, val int) {
	s.sha.mat.dat[ind] += float32(val)
}

func (s *setter) Sta(rat float32) {
	var hal int
	{
		hal = len(s.sha.sta.dat) / 2
	}

	var fac float32
	{
		fac = 1 / float32(hal)
	}

	if rat < 0 {
		s.sha.sta.dat[int(rat/-fac)]++
	}

	if rat == 0 {
		s.sha.sta.dat[hal]++
	}

	if rat > 0 {
		s.sha.sta.dat[int(rat/+fac)]++
	}
}
