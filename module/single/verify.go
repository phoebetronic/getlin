package single

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/metric"
)

func (m *Module) Verify(vec [][2]getlin.Vector) getlin.Metric {
	var met getlin.Metric
	{
		met = metric.New(metric.Config{})
	}

	for _, x := range vec {
		for _, y := range x {
			var t float32
			{
				t = y.Out()[0]
			}

			var p float32
			{
				p = m.Search(y).Out()[0]
			}

			{
				met.Set().Err(t, p)
			}
		}
	}

	return met
}
