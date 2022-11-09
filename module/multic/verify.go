package multic

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/metric"
	"github.com/phoebetron/getlin/stat32"
)

func (m *Module) Verify(vec [][2]getlin.Vector) getlin.Metric {
	var met getlin.Metric
	{
		met = metric.New(metric.Config{})
	}

	for _, x := range vec {
		for _, y := range x {
			var t int
			{
				t = y.Cla()
			}

			var p int
			{
				p = stat32.Argmax(m.Search(y).Out())
			}

			{
				met.Set().Err(float32(t), float32(p))
			}
		}
	}

	return met
}
