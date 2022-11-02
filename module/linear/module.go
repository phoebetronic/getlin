package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/active/harrel"
	"github.com/phoebetron/getlin/clause"
	"github.com/phoebetron/getlin/metric"
	"github.com/phoebetron/getlin/serial"
	"github.com/phoebetron/getlin/shaper"
)

type Module struct {
	cla []getlin.Clause
	met getlin.Metric
	ser serial.Interface
	sha getlin.Shaper
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

	var inp int
	var out int
	var sta int
	{
		inp = con.Inp
		out = con.Out
		sta = (2 * con.Inp) * con.Out
	}

	return &Module{
		cla: cla,
		met: met,
		ser: con.Ser,
		sha: shaper.New(shaper.Config{Inp: inp, Out: out, Sta: sta}),
	}
}
