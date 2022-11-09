package vector

type Config struct {
	// Cla provides optional class identifier information for different Module
	// implementations to leverage internally.
	Cla int
	// Inp carries the configured feature weights of this Vector.
	Inp []float32
	// Out carries the true weights or weight predictions of this Vector.
	Out []float32
}

func (c Config) Verify() {}
