package metric

type setter struct {
	sha *shared
}

func (s *setter) Err(tru float32, prd float32) {
	s.sha.err.cou += 1
	s.sha.err.los += absf32(tru - prd)
}

func (s *setter) Sta(ind int, val int) {
	s.sha.sta.dat[ind] += val
}

func absf32(f float32) float32 {
	if f < 0 {
		return -f
	}

	return f
}
