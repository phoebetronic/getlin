package vector

type Status struct {
	ini bool
	raw bool
}

func (s *Status) Fai() bool {
	return !s.raw
}

func (s *Status) Ini() bool {
	return s.ini
}

func (s *Status) Suc() bool {
	return s.raw
}

func (s *Status) Upd(sta bool) {
	s.ini = true
	s.raw = sta
}
