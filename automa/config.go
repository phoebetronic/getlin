package automa

import "math/rand"

type Config struct {
	// Poi provides the option to initialize the internal state pointer to the
	// value given here. For now this is mostly used for testing.
	Poi int
	// Sta is the number of states S along a single side of the states
	// distribution, where the final states distribution will be 2 * S. Below is
	// illustrated a state distribution 8 = S = 4.
	//
	//     1    2    3    4    |    5    6    7    8
	//
	Sta int
}

func (c Config) Ensure() Config {
	if c.Poi == 0 {
		c.Poi = rand.Intn(c.Sta+1-c.Sta+1) + c.Sta // TODO configure real random
	}

	return c
}

func (c Config) Verify() {
	if c.Sta == 0 {
		panic("Config.Sta must not be empty")
	}
}
