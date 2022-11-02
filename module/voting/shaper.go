package voting

import "github.com/phoebetron/getlin"

func (m *Module) Shaper() getlin.Shaper {
	return m.sha
}
