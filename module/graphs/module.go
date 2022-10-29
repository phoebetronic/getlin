package graphs

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/cacher"
	"github.com/phoebetron/getlin/mapper"
)

type Module struct {
	cac getlin.Cacher
	mod [][]getlin.Module
	mpr getlin.Mapper
}

func New(con Config) *Module {
	{
		con.Verify()
	}

	var mod [][]getlin.Module
	for _, x := range con.Mod {
		var lay []getlin.Module

		for _, y := range x {
			lay = append(lay, y)
		}

		{
			mod = append(mod, lay)
		}
	}

	return &Module{
		cac: cacher.New(cacher.Config{}),
		mod: con.Mod,
		mpr: mapper.New(mapper.Config{Mod: mod}),
	}
}
