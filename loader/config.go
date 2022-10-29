package loader

type Config struct {
	Git bool
	Org string
	Rep string
}

func (c *Config) Verify() {
	if !c.Git {
		panic("Config.Git must not be empty")
	}
	if c.Org == "" {
		panic("Config.Org must not be empty")
	}
	if c.Rep == "" {
		panic("Config.Rep must not be empty")
	}
}
