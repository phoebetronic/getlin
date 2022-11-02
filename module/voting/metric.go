package voting

import "github.com/phoebetron/getlin"

func (m *Module) Metric() getlin.Metric {
	if m.wei != nil {
		return m.wei.Metric()
	}

	return nil
}
