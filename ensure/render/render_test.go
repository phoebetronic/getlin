package render

import (
	"flag"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
	"github.com/phoebetron/getlin/module/graphs"
	"github.com/phoebetron/getlin/module/linear"
	"github.com/phoebetron/getlin/module/voting"
)

var upd = flag.Bool("u", false, "whether to update golden files")

func musgra(mod [][]getlin.Module) getlin.Module {
	return graphs.New(graphs.Config{
		Mpr: mapper.New(mapper.Config{Mod: mod}),
	})
}

func muslin(inp int, out int) getlin.Module {
	return linear.New(linear.Config{Inp: inp, Out: out, Sta: 4})
}

func musvot(inp int) getlin.Module {
	return voting.New(voting.Config{Inp: inp})
}
