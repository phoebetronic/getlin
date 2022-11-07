package multic

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/shaper"
)

// Module implements the multi class version of the Tsetlin Machine, voting for
// a number of binary prediction result.
type Module struct {
	sha getlin.Shaper
	sin []getlin.Module
}

func New(con Config) *Module {
	{
		con.Verify()
	}

	return &Module{
		sha: shaper.New(shaper.Config{
			Cla: con.Sin[0].Shaper().Cla() * len(con.Sin),
			Inp: con.Sin[0].Shaper().Inp(),
			Out: len(con.Sin),
			Sta: con.Sin[0].Shaper().Sta() * len(con.Sin),
		}),
		sin: con.Sin,
	}
}
