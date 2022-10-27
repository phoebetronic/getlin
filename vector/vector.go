package vector

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

func (v *Vector) And() Interface {
	var vec Interface
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

func (v *Vector) Len() int {
	return len(v.bit)
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

func (v *Vector) Tru(ind int) bool {
	return v.tru[ind]
}

func (v *Vector) Zer() bool {
	return v.zer && !v.one
}
