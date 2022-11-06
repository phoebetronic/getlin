package clause

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/automa"
)

type Clause struct {
	fre float32
	neg []getlin.Automa
	pos []getlin.Automa
	ran getlin.Random
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
		fre: con.Fre,
		neg: neg,
		pos: pos,
		ran: con.Ran,
	}
}
