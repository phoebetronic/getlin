package mapper

import (
	"github.com/phoebetron/getlin"
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

	var lay int
	for _, x := range con.Mod[0] {
		lay += x.Clause()
	}

	var pre []getlin.Module
	for i, x := range con.Mod {
		var out int
		for _, y := range x {
			out += y.Clause()
		}

		if out != lay {
			con.Err(tracer.Maskf(
				invalidMappingError,
				"the cumulative output Vector length %d defined by layer 0 is incompatible with the output Vector length %d defined by layer %d",
				lay,
				out,
				i,
			))
		}

		for _, y := range x {
			if y.Clause() != 0 && i > 0 {
				var inp int
				{
					inp = y.Automa() / 2 / y.Clause()
				}

				if inp != lay {
					con.Err(tracer.Maskf(
						invalidMappingError,
						"the cumulative output Vector length %d defined by layer %d is incompatible with the input Vector length %d defined by Module %d in layer %d",
						lay,
						i-1,
						inp,
						len(lin),
						i,
					))
				}
			}

			var a Linker
			{
				a = lin[y]
			}

			{
				a.Ind = len(lin)
				a.Lay = i
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

	for _, x := range con.Mod {
		var lef float32

		for _, y := range x {
			if y.Clause() == 0 {
				continue
			}

			var rig float32
			{
				rig = lef + float32(y.Clause())
			}

			var l Linker
			{
				l = lin[y]
			}

			{
				l.Tru = []int{int(lef), int(rig)}
			}

			{
				lin[y] = l
			}

			{
				lef = rig
			}
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

func (m *Mapper) Tru(mod getlin.Module) []int {
	return m.lin[mod].Tru
}
