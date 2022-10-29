package serial

import "github.com/phoebetron/getlin"

type Interface interface {
	Reader([]byte) getlin.Vector
	Writer(getlin.Vector) []byte
}
