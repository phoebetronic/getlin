package normal

import "github.com/phoebetron/getlin"

type and struct{}

func And() Interface {
	return &and{}
}

func (s *and) Normal(vec getlin.Vector) getlin.Vector {
	return vec.And()
}
