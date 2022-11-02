package vector

import "github.com/phoebetron/getlin"

type binary struct {
	raw []bool

	one float32
	zer float32
}

func newbin(raw ...bool) *binary {
	var one float32
	var zer float32
	for _, x := range raw {
		if x {
			one++
		} else {
			zer++
		}
	}

	return &binary{
		raw: raw,

		one: one,
		zer: zer,
	}
}

func (b *binary) Add(bit bool) {
	if bit {
		b.one++
	} else {
		b.zer++
	}

	{
		b.raw = append(b.raw, bit)
	}
}

func (b *binary) And() getlin.Binary {
	if b.One() {
		return newbin(true)
	}

	return newbin(false)
}

func (b *binary) Cop() getlin.Binary {
	return newbin(b.raw...)
}

func (b *binary) Eql(bin getlin.Binary) bool {
	return eql(b.Raw(), bin.Raw())
}

func (b *binary) Maj() bool {
	return b.one-b.zer >= 1
}

func (b *binary) Neg(ind int) bool {
	return !b.raw[ind]
}

func (b *binary) One() bool {
	return b.one != 0 && b.zer == 0
}

func (b *binary) Pos(ind int) bool {
	return b.raw[ind]
}

func (b *binary) Raw() []bool {
	return append([]bool{}, b.raw...)
}

func (b *binary) Spl(ind int) (getlin.Binary, getlin.Binary) {
	var lef []bool
	var rig []bool
	if len(b.raw) > ind {
		{
			lef = b.raw[:ind]
			rig = b.raw[ind:]
		}
	} else if len(b.raw) != 0 {
		lef = b.raw
	}

	return newbin(lef...), newbin(rig...)
}

func (b *binary) Wei(bit bool) float32 {
	if bit {
		return b.one / (b.zer + b.one)
	}

	return b.zer / (b.zer + b.one)
}

func (b *binary) Zer() bool {
	return b.zer != 0 && b.one == 0
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
