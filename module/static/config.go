package static

import "github.com/phoebetron/getlin"

type Config struct {
	Out getlin.Vector
}

func (c Config) Verify() {
	if c.Out == nil {
		panic("Config.Out must not be empty")
	}
}
