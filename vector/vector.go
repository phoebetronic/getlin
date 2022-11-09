package vector

type Vector struct {
	cla int
	inp []float32
	out []float32
}

func New(con Config) *Vector {
	{
		con.Verify()
	}

	return &Vector{
		cla: con.Cla,
		inp: con.Inp,
		out: con.Out,
	}
}

func (v *Vector) Cla() int {
	return v.cla
}

func (v *Vector) Inp() []float32 {
	return v.inp
}

func (v *Vector) Out() []float32 {
	return v.out
}
