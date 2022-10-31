package static

type Config struct {
	F32 float32
}

func (c Config) Verify() {
	if c.F32 < 0 {
		panic("Config.F32 must not be smaller than 0")
	}
	if c.F32 > 1 {
		panic("Config.F32 must not be greater than 1")
	}
}
