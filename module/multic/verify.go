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

		var t []float32
		{
			t = make([]float32, len(o))
		}

		{
			t[c] = 1
		}

		var p []float32
		{
			p = make([]float32, len(o))
		}

		{
			p[a] = 1
		}

		for i := range o {
			met.Set().Err(t[i], p[i])
		}
	}

	return met
}
