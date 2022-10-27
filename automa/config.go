package automa

type Config struct {
	// Sta is the number of states S along a single side of the states
	// distribution, where the final states distribution will be (2 * S) + 1.
	// Below is illustrated a state distribution 9 = S = 4.
	//
	//     -    S    3    2    1    0    1    2    3    S    +
	//
	Sta int
}

func (c *Config) Verify() {
	if c.Sta == 0 {
		panic("Config.Sta must not be empty")
	}
}
