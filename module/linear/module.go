package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/active"
	"github.com/phoebetron/getlin/clause"
	"github.com/phoebetron/getlin/metric"
	"github.com/phoebetron/getlin/serial"
)

type Module struct {
	cla []getlin.Clause
	met getlin.Metric
	ser serial.Interface
}

func New(con Config) *Module {
	{
		con.Verify()
	}

	var act active.Interface
	{
		act = active.Hrl()
	}

	var met getlin.Metric
	{
		met = metric.New(metric.Config{})
	}

	var cla []getlin.Clause
	for i := 0; i < con.Out; i++ {
		cla = append(cla, clause.New(clause.Config{
			Act: act,
			Met: met,
			Ser: con.Ser,
			Sta: con.Sta,
			Tas: con.Inp,
		}))
	}

	return &Module{
		cla: cla,
		met: met,
		ser: con.Ser,
	}
}
