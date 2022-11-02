package graphs

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/cacher"
	"github.com/phoebetron/getlin/shaper"
)

type Module struct {
	cac getlin.Cacher
	mpr getlin.Mapper
	sha getlin.Shaper
}

func New(con Config) *Module {
	{
		con.Verify()
	}

	var inp int
	{
		inp = con.Mpr.All()[0][0].Shaper().Inp()
	}

	var out int
	for _, x := range con.Mpr.All()[len(con.Mpr.All())-1] {
		out += x.Shaper().Out()
	}

	var sta int
	for _, x := range con.Mpr.All() {
		for _, y := range x {
			sta += y.Shaper().Sta()
		}
	}

	return &Module{
		cac: cacher.New(cacher.Config{}),
		mpr: con.Mpr,
		sha: shaper.New(shaper.Config{Inp: inp, Out: out, Sta: sta}),
	}
}
