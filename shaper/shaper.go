package shaper

type Shaper struct {
	inp int
	out int
	sta int
}

func New(con Config) *Shaper {
	{
		con.Verify()
	}

	return &Shaper{
		inp: con.Inp,
		out: con.Out,
		sta: con.Sta,
	}
}

func (s *Shaper) Inp() int {
	return s.inp
}

func (s *Shaper) Out() int {
	return s.out
}

func (s *Shaper) Sta() int {
	return s.sta
}
