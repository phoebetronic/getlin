package graphs

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
	"github.com/phoebetron/getlin/module/graphs"
	"github.com/phoebetron/getlin/module/linear"
	"github.com/phoebetron/getlin/module/voting"
	"github.com/phoebetron/getlin/random/static"
	"github.com/phoebetron/getlin/shaper"
	"github.com/phoebetron/getlin/vector"
)

type tester struct {
	Sea getlin.Vector
	Sha getlin.Shaper
	Upd getlin.Vector
}

func (s *tester) Mapper() getlin.Mapper {
	return nil
}

func (s *tester) Metric() getlin.Metric {
	return nil
}

func (s *tester) Search(vec getlin.Vector) getlin.Vector {
	return s.Sea
}

func (s *tester) Shaper() getlin.Shaper {
	return s.Sha
}

func (s *tester) States() []float32 {
	return nil
}

func (s *tester) Update(vec getlin.Vector) {
	s.Upd = vec
}

func musgra(mod [][]getlin.Module) getlin.Module {
	return graphs.New(graphs.Config{
		Mpr: mapper.New(mapper.Config{Mod: mod}),
	})
}

func muslin(inp int, out int) getlin.Module {
	return linear.New(linear.Config{
		Inp: inp,
		Out: out,
		Ran: static.New(static.Config{F32: 0.05}),
		Sta: 4,
	})
}

func mussha(inp int, out int) getlin.Shaper {
	return shaper.New(shaper.Config{
		Inp: inp,
		Out: out,
	})
}

func musvec(inp []int, out []int) getlin.Vector {
	return vector.New(vector.Config{
		Inp: vector.ToBool(inp...),
		Out: vector.ToBool(out...),
	})
}

func musvot(inp int) getlin.Module {
	return voting.New(voting.Config{Inp: inp})
}
