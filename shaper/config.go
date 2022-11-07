package shaper

type Config struct {
	Cla int
	Inp int
	Out int
	Sta int
}

func (c Config) Verify() {
	if c.Cla == 0 {
		panic("Config.Cla must not be empty")
	}
	if c.Inp == 0 {
		panic("Config.Inp must not be empty")
	}
	if c.Out == 0 {
		panic("Config.Out must not be empty")
	}
	if c.Sta == 0 {
		panic("Config.Sta must not be empty")
	}
}
