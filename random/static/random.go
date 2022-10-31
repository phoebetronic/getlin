package static

type Random struct {
	f32 float32
}

func New(con Config) *Random {
	{
		con.Verify()
	}

	return &Random{
		f32: con.F32,
	}
}

func (r *Random) F32() float32 {
	return r.f32
}
