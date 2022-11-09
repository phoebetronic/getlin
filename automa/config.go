package automa

import "math/rand"

type Config struct {
	// Poi provides the option to initialize the internal state pointer to the
	// value given here. For now this is mostly used for testing.
	Poi float32
}

func (c Config) Ensure() Config {
	if c.Poi == 0 {
		c.Poi = rand.Float32() // TODO configure real random
	}

	return c
}

func (c Config) Verify() {}
