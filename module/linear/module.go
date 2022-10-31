package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/active/harrel"
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
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	var act getlin.Active
	{
		act = harrel.New(harrel.Config{})
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
			Ran: con.Ran,
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
