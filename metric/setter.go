package metric

import "github.com/phoebetron/getlin/stat32"

type setter struct {
	sha *shared
}

func (s *setter) Err(tru float32, prd float32) {
	s.sha.err.cou += 1
	s.sha.err.los += stat32.Absf32(tru - prd)
}

func (s *setter) Sta(ind int, val int) {
	s.sha.sta.dat[ind] += val
}
