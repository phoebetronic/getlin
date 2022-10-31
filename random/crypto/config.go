package crypto

import (
	"math/rand"
)

type Config struct {
	Src rand.Source
}

func (c Config) Ensure() Config {
	if c.Src == nil {
		c.Src = source{}
	}

	return c
}

func (c Config) Verify() {
	if c.Src == nil {
		panic("Config.Src must not be empty")
	}
}
