package graphs

import "github.com/phoebetron/getlin"

func (m *Module) Metric() getlin.Metric {
	var met getlin.Metric

	for _, x := range m.mpr.All() {
		for _, y := range x {
			for i, d := range y.Metric().Get().Mat().Raw() {
				met.Set().Mat(i, d)
			}
			for i, d := range y.Metric().Get().Sta().Raw() {
				met.Set().Sta(i, d)
			}
		}
	}

	return met
}
