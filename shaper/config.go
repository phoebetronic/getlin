package shaper

type Config struct {
	Inp int
	Out int
	Sta int
}

func (c Config) Verify() {
	if c.Inp == 0 {
		panic("Config.Inp must not be empty")
	}
	if c.Out == 0 {
		panic("Config.Out must not be empty")
	}
}
