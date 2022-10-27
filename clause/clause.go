package clause

import (
	"github.com/phoebetron/getlin/active"
	"github.com/phoebetron/getlin/automa"
	"github.com/phoebetron/getlin/metric"
	"github.com/phoebetron/getlin/serial"
)

type Clause struct {
	act active.Interface
	met metric.Interface
	ser serial.Interface

	neg []automa.Interface
	pos []automa.Interface
}

func New(con Config) *Clause {
	{
		con.Verify()
	}

	var neg []automa.Interface
	var pos []automa.Interface
	for i := 0; i < con.Tas; i++ {
		neg = append(neg, automa.New(automa.Config{Sta: con.Sta}))
		pos = append(pos, automa.New(automa.Config{Sta: con.Sta}))
	}

	return &Clause{
		act: con.Act,
		met: con.Met,
		ser: con.Ser,

		neg: neg,
		pos: pos,
	}
}
