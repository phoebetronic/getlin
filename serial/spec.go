package serial

import "github.com/phoebetron/getlin/vector"

type Interface interface {
	Reader([]byte) vector.Interface
	Writer(vector.Interface) []byte
}
