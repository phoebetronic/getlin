package graphs

import "github.com/phoebetron/getlin"

type Config struct {
	// Mpr manages the graph of Modules representing this Graphs Module's
	// architecture. See mapper.Config for structural requirements.
	Mpr getlin.Mapper
}

func (c Config) Verify() {
	if c.Mpr == nil {
		panic("Config.Mpr must not be empty")
	}
}
