package graphs

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/cacher"
	"github.com/phoebetron/getlin/mapper"
)

type Module struct {
	cac getlin.Cacher
	mpr getlin.Mapper
}

func New(con Config) *Module {
	{
		con.Verify()
	}

	return &Module{
		cac: cacher.New(cacher.Config{}),
		mpr: mapper.New(mapper.Config{Mod: con.Mod}),
	}
}
