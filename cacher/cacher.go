package cacher

import "github.com/phoebetron/getlin"

type Cacher struct {
	ind int
	vec map[int]getlin.Linker
}

func New(con Config) *Cacher {
	{
		con.Verify()
	}

	return &Cacher{
		vec: map[int]getlin.Linker{},
	}
}

func (m *Cacher) Create(inp getlin.Vector) {
	m.ind++
	m.vec[m.ind] = getlin.Linker{Inp: inp}
}

func (m *Cacher) Delete() {
	m.ind = 0
	m.vec = map[int]getlin.Linker{}
}

func (m *Cacher) Latest() getlin.Linker {
	return m.vec[m.ind]
}

func (m *Cacher) Update(out getlin.Vector) {
	vec := m.vec[m.ind]
	vec.Out = out
	m.vec[m.ind] = vec
}
