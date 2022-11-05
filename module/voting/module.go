package voting

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/clause"
	"github.com/phoebetron/getlin/shaper"
)

type Module struct {
	cla []getlin.Clause
	ran getlin.Random
	sha getlin.Shaper
	thr float32
}

func New(con Config) *Module {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	var cla []getlin.Clause
	for i := 0; i < con.Cla; i++ {
		cla = append(cla, clause.New(clause.Config{
			Fre: con.Fre,
			Ran: con.Ran,
			Sta: con.Sta,
			Tas: con.Inp,
		}))
	}

	var inp int
	var out int
	var sta int
	{
		inp = con.Inp
		out = con.Cla
		sta = (2 * con.Inp) * con.Cla
	}

	return &Module{
		cla: cla,
		ran: con.Ran,
		sha: shaper.New(shaper.Config{Inp: inp, Out: out, Sta: sta}),
		thr: float32(con.Thr),
	}
}
