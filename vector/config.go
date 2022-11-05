package vector

import "github.com/phoebetron/getlin"

type Config struct {
	// Ind is an optional instruction of Module update paths at the root of a
	// Graphs Module. Considering the example below, feedback would only be
	// given to the Modules at Module index 2 and 8, in the first and second
	// layer respectively.
	//
	//     [
	//         [0 0 1 0 0 0 0 0 1 0]
	//         [0 0 1 0 0 0 0 0 1 0]
	//     ]
	//
	Ind [][]bool
	// Inp carries the input bits consumed by any given Module, be it inference
	// or training.
	Inp []uint8
	// Out provides either the externally defined true labels or carries the
	// predicted labels.
	Out []uint8
	// Sta is the internally managed Vector status and should never be
	// manipulated from the outside.
	Sta getlin.Status
}

func (c Config) Ensure() Config {
	if c.Sta == nil {
		c.Sta = &Status{}
	}

	return c
}

func (c Config) Verify() {
}
