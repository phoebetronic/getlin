package mapper

import "github.com/phoebetron/getlin"

type Mapper struct {
	all [][]getlin.Module
	lin map[getlin.Module]Linker
}

func New(con Config) *Mapper {
	{
		con.Verify()
	}

	lin := map[getlin.Module]Linker{}
	pre := []getlin.Module{}

	for i := 0; i < len(con.Mod); i++ {
		var lay []getlin.Module
		{
			lay = con.Mod[i]
		}

		for j := 0; j < len(lay); j++ {
			var c getlin.Module
			{
				c = lay[j]
			}

			var a Linker
			{
				a = lin[c]
			}

			{
				a.Ind = len(lin)
			}

			for _, p := range pre {
				var b Linker
				{
					b = lin[p]
				}

				{
					a.Abo = append(a.Abo, p)
					b.Bel = append(b.Bel, c)
				}

				{
					lin[p] = b
				}
			}

			{
				lin[c] = a
			}
		}

		{
			pre = lay
		}
	}

	return &Mapper{
		all: con.Mod,
		lin: lin,
	}
}

func (m *Mapper) Abo(mod getlin.Module) []getlin.Module {
	return m.lin[mod].Abo
}

func (m *Mapper) All() [][]getlin.Module {
	return m.all
}

func (m *Mapper) Bel(mod getlin.Module) []getlin.Module {
	return m.lin[mod].Bel
}

func (m *Mapper) Ind(mod getlin.Module) int {
	return m.lin[mod].Ind
}
