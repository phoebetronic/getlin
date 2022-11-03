package mapper

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
	"github.com/xh3b4sd/tracer"
)

type Mapper struct {
	lin map[getlin.Module]Linker
	mod [][]getlin.Module
}

func New(con Config) *Mapper {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	lin := map[getlin.Module]Linker{}

	var fir int
	for _, x := range con.Mod[0] {
		fir += x.Shaper().Out()
	}

	var out int
	for _, x := range con.Mod[len(con.Mod)-1] {
		out += x.Shaper().Out()
	}

	var pre []getlin.Module
	for i, x := range con.Mod {
		var lay []int
		var sum int
		for _, y := range x {
			lay = append(lay, y.Shaper().Out())
			sum += y.Shaper().Out()
		}

		if sum != fir && sum/len(x) != 1 {
			con.Err(tracer.Maskf(
				invalidMappingError,
				"the cumulative output Vector length %d defined by layer 0 is incompatible with the output Vector length %d defined by layer %d",
				fir,
				sum,
				i,
			))
		}

		for j, y := range x {
			if i > 0 && y.Shaper().Inp() != fir {
				con.Err(tracer.Maskf(
					invalidMappingError,
					"the cumulative output Vector length %d defined by layer %d is incompatible with the input Vector length %d defined by Module %d in layer %d",
					fir,
					i-1,
					y.Shaper().Inp(),
					len(lin),
					i,
				))
			}

			var a Linker
			{
				a = lin[y]
			}

			var tru [2]int
			{
				tru = vector.Labels(lay, out, j)
			}

			if tru[1] == 0 {
				con.Err(tracer.Maskf(
					invalidMappingError,
					"Module %d in layer %d cannot receive accurate feedback considering its position, output Vector length and respective true label relationship",
					len(lin),
					i,
				))
			}

			{
				a.Ind = len(lin)
				a.Lay = i
				a.Tru = tru
			}

			for _, p := range pre {
				var b Linker
				{
					b = lin[p]
				}

				{
					a.Abo = append(a.Abo, p)
					b.Bel = append(b.Bel, y)
				}

				{
					lin[p] = b
				}
			}

			{
				lin[y] = a
			}
		}

		{
			pre = x
		}
	}

	return &Mapper{
		lin: lin,
		mod: con.Mod,
	}
}

func (m *Mapper) Abo(mod getlin.Module) []getlin.Module {
	return m.lin[mod].Abo
}

func (m *Mapper) All() [][]getlin.Module {
	return m.mod
}

func (m *Mapper) Bel(mod getlin.Module) []getlin.Module {
	return m.lin[mod].Bel
}

func (m *Mapper) Ind(mod getlin.Module) int {
	return m.lin[mod].Ind
}

func (m *Mapper) Lay(mod getlin.Module) int {
	return m.lin[mod].Lay
}

func (m *Mapper) Tru(mod getlin.Module) [2]int {
	return m.lin[mod].Tru
}
