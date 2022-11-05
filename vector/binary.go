package vector

import (
	"bytes"

	"github.com/phoebetron/getlin"
)

type binary struct {
	raw []uint8

	one float32
	zer float32
}

func newbin(raw ...uint8) *binary {
	var one float32
	var zer float32
	for _, x := range raw {
		if x == 1 {
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

func (b *binary) Add(bit uint8) {
	if bit == 1 {
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
		return newbin(1)
	}

	return newbin(0)
}

func (b *binary) Cop() getlin.Binary {
	return newbin(b.raw...)
}

func (b *binary) Cou(bit uint8) int {
	if bit == 1 {
		return int(b.one)
	}

	return int(b.zer)
}

func (b *binary) Eql(bin getlin.Binary) bool {
	return bytes.Equal(b.Raw(), bin.Raw())
}

func (b *binary) Maj() uint8 {
	if b.one-b.zer >= 1 {
		return 1
	}

	return 0
}

func (b *binary) Neg(ind int) uint8 {
	return ^b.raw[ind]
}

func (b *binary) One() bool {
	return b.one != 0 && b.zer == 0
}

func (b *binary) Pos(ind int) uint8 {
	return b.raw[ind]
}

func (b *binary) Raw() []uint8 {
	return append([]uint8{}, b.raw...)
}

func (b *binary) Wei(bit uint8) float32 {
	if bit == 1 {
		return b.one / (b.zer + b.one)
	}

	return b.zer / (b.zer + b.one)
}

func (b *binary) Zer() bool {
	return b.zer != 0 && b.one == 0
}
