package normal

import (
	"github.com/phoebetron/getlin/vector"
)

type and struct{}

func And() Interface {
	return &and{}
}

func (s *and) Normal(vec vector.Interface) vector.Interface {
	return vec.And()
}
