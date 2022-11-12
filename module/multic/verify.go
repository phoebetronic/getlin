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
		var c int
		{
			c = x[1].Cla()
		}

		var o []float32
		{
			o = m.Search(x[1]).Out()
		}

		var a int
		{
			a = stat32.Argmax(o)
		}

		if o[a] == 0 {
			met.Set().Err(1, 0)
		} else {
			if c == a {
				met.Set().Err(1, 1)
			}
			if c != a {
				met.Set().Err(1, o[a])
			}
		}
	}

	return met
}
