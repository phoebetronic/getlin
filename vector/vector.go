package vector

import "github.com/phoebetron/getlin"

type Vector struct {
	ind [][]bool
	inp *binary
	out *binary
	sta getlin.Status
}

func New(con Config) *Vector {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	return &Vector{
		ind: con.Ind,
		inp: newbin(con.Inp...),
		out: newbin(con.Out...),
		sta: con.Sta,
	}
}

func (v *Vector) Cop() getlin.Vector {
	return New(Config{Inp: v.Inp().Raw(), Out: v.Inp().Raw()})
}

func (v *Vector) Eql(vec getlin.Vector) bool {
	return v.Inp().Eql(vec.Inp()) && v.Out().Eql(vec.Out())
}

func (v *Vector) Ind() [][]bool {
	return v.ind
}

func (v *Vector) Inp() getlin.Binary {
	return v.inp
}

func (v *Vector) Out() getlin.Binary {
	return v.out
}

func (v *Vector) Sta() getlin.Status {
	return v.sta
}
