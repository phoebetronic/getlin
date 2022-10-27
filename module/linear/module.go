package linear

import (
	"github.com/phoebetron/getlin/active"
	"github.com/phoebetron/getlin/clause"
	"github.com/phoebetron/getlin/metric"
	"github.com/phoebetron/getlin/serial"
)

type Module struct {
	cla []clause.Interface
	met metric.Interface
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

	var met metric.Interface
	{
		met = metric.New(metric.Config{})
	}

	var cla []clause.Interface
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
