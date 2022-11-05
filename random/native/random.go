package native

import "math/rand"

type Random struct {
}

func New(con Config) *Random {
	{
		con.Verify()
	}

	return &Random{}
}

func (r *Random) F32() float32 {
	return rand.Float32()
}
