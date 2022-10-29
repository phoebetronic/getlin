package mapper

import "github.com/phoebetron/getlin"

type Config struct {
	Mod [][]getlin.Module
}

func (c *Config) Verify() {
	if len(c.Mod) == 0 {
		panic("Config.Mod must not be empty")
	}
}
