package graphs

import "github.com/phoebetron/getlin"

type Config struct {
	// Mod is the graph of modules representing this Graphs Module's
	// architecture. See mapper.Config for structural requirements.
	Mod [][]getlin.Module
}

func (c Config) Verify() {
	if len(c.Mod) == 0 {
		panic("Config.Mod must not be empty")
	}
}
