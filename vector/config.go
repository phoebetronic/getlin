package vector

import "github.com/phoebetron/getlin"

type Config struct {
	Inp []bool
	Out []bool
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
