package single

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/clause"
	"github.com/phoebetron/getlin/shaper"
)

// Module implements the classic version of the Tsetlin Machine, voting for a
// single binary prediction result.
type Module struct {
	cla []getlin.Clause
	ran getlin.Random
	sha getlin.Shaper
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
			Tas: con.Inp,
		}))
	}

	return &Module{
		cla: cla,
		ran: con.Ran,
		sha: shaper.New(shaper.Config{
			Cla: con.Cla,
			Inp: con.Inp,
			Out: 1,
			Sta: (2 * con.Inp) * con.Cla,
		}),
	}
}
