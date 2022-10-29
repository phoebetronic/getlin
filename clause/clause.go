package clause

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/active"
	"github.com/phoebetron/getlin/automa"
	"github.com/phoebetron/getlin/serial"
)

type Clause struct {
	act active.Interface
	met getlin.Metric
	ser serial.Interface

	neg []getlin.Automa
	pos []getlin.Automa
}

func New(con Config) *Clause {
	{
		con.Verify()
	}

	var neg []getlin.Automa
	var pos []getlin.Automa
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
