package voting

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/shaper"
)

type Module struct {
	sha getlin.Shaper
	wei getlin.Module
}

func New(con Config) *Module {
	{
		con.Verify()
	}

	var inp int
	{
		inp = con.Inp
	}

	var out int
	{
		out = 1
	}

	var sta int
	if con.Wei != nil {
		sta = con.Wei.Shaper().Sta()
	}

	return &Module{
		sha: shaper.New(shaper.Config{Inp: inp, Out: out, Sta: sta}),
		wei: con.Wei,
	}
}
