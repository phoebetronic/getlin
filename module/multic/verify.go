package multic

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
			var c int
			var t float32
			{
				c = y.Cla()
				t = y.Out()[0]
			}

			var p []float32
			{
				p = m.Search(y).Out()
			}

			{
				met.Set().Err(t, p[c])
			}
		}
	}

	return met
}
