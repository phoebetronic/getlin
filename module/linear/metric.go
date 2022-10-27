package linear

import "github.com/phoebetron/getlin/metric"

func (m *Module) Metric() metric.Interface {
	return m.met
}
