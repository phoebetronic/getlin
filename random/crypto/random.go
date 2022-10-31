package crypto

import (
	"math/rand"
)

type Random struct {
	ran *rand.Rand
}

func New(con Config) *Random {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	return &Random{
		ran: rand.New(con.Src),
	}
}

func (r *Random) F32() float32 {
	return r.ran.Float32()
}
