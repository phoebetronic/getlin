package mapper

import (
	"github.com/phoebetron/getlin"
)

type Config struct {
	// Err is an optimal error handler which panics by default. Errors forwarded
	// here are only produced within the constructor function New. Any errors
	// raised indicate an invalids setting for the graph architecture. Err is
	// mainly used for testing.
	Err func(error)
	// Mod represents the graph of modules to be mapped. The mapper validates
	// the graph structure and provides lookup primitives of associative
	// features inherent to the modules comprising any given architecture.
	// Information like Module relationships across layers and true label index
	// ranges per Module can be accessed using a successfully generated mapping.
	// Two requirements restrict the graph composition upon instantiation.
	//
	//     1. All layers must produce the same cumulative output Vector length.
	//
	//     2. All modules must receive an input Vector of length compatible with
	//     the cumulative output Vector length produced by their preceding
	//     layers.
	//
	Mod [][]getlin.Module
}

func (c Config) Ensure() Config {
	if c.Err == nil {
		c.Err = func(err error) {
			panic(err)
		}
	}

	return c
}

func (c Config) Verify() {
	if c.Err == nil {
		panic("Config.Err must not be empty")
	}
	if len(c.Mod) == 0 {
		panic("Config.Mod must not be empty")
	}
}
