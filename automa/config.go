package automa

import "math/rand"

type Config struct {
	// Poi provides the option to initialize the internal state pointer to the
	// value given here. For now this is mostly used for testing.
	Poi float32
}

func (c Config) Ensure() Config {
	if c.Poi == 0 {
		if rand.Float32() <= 0.5 {
			c.Poi = 0.5
		} else {
			c.Poi = 0.6
		}
	}

	return c
}

func (c Config) Verify() {}
