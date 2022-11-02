package voting

import "github.com/phoebetron/getlin"

type Config struct {
	// Inp defines the input Vector length this Module accepts during Search and
	// Update.
	Inp int
	// Wei is the optional weights Module learning a desired voting
	// distribution. If Wei is empty a simple majority voting takes place. That
	// is, if the given input Vector were to represent [0 0 1], then 0 would be
	// returned as majority voting result.
	Wei getlin.Module
}

func (c Config) Verify() {
	if c.Inp == 0 {
		panic("Config.Inp must not be empty")
	}
}
