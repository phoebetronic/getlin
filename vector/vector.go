package vector

import "github.com/phoebetron/getlin"

type Vector struct {
	bit []bool
	tru []bool

	one bool
	zer bool
}

func New(con Config) *Vector {
	{
		con.Verify()
	}

	var one bool
	var zer bool
	for _, x := range con.Bit {
		if x {
			one = true
		}
		if !x {
			zer = true
		}
	}

	return &Vector{
		bit: con.Bit,
		tru: con.Tru,

		one: one,
		zer: zer,
	}
}

func (v *Vector) Add(bit bool) {
	if bit {
		v.one = true
	} else {
		v.zer = true
	}

	{
		v.bit = append(v.bit, bit)
	}
}

func (v *Vector) And() getlin.Vector {
	var vec getlin.Vector
	{
		vec = New(Config{})
	}

	if v.One() {
		vec.Add(true)
	} else {
		vec.Add(false)
	}

	return vec
}

func (v *Vector) Bit() []bool {
	return append([]bool{}, v.bit...)
}

func (v *Vector) Eql(vec getlin.Vector) bool {
	return eql(v.Bit(), vec.Bit()) && eql(v.Tru(), vec.Tru())
}

func (v *Vector) Neg(ind int) bool {
	return !v.bit[ind]
}

func (v *Vector) One() bool {
	return v.one && !v.zer
}

func (v *Vector) Pos(ind int) bool {
	return v.bit[ind]
}

func (v *Vector) Tru() []bool {
	return append([]bool{}, v.tru...)
}

func (v *Vector) Zer() bool {
	return v.zer && !v.one
}

func eql(a []bool, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
