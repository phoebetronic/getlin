package metric

type setter struct {
	sha *shared
}

func (s *setter) Mat(ind int, val int) {
	s.sha.mat.dat[ind] += val
}

func (s *setter) Sta(ind int, val int) {
	s.sha.sta.dat[ind] += val
}
