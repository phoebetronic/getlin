package crypto

import (
	"crypto/rand"
	"encoding/binary"
)

type source struct{}

func (s source) Seed(seed int64) {}

func (s source) Int63() int64 {
	return int64(s.uint64() & ^uint64(1<<63))
}

func (s source) uint64() uint64 {
	var v uint64
	{
		err := binary.Read(rand.Reader, binary.BigEndian, &v)
		if err != nil {
			panic(err)
		}
	}

	return v
}
