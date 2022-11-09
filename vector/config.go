package vector

type Config struct {
	// Cla provides optional class identifier information for different Module
	// implementations to leverage internally.
	Cla int
	// Inp carries the input bits consumed by any given Module, be it inference
	// or training.
	Inp []float32
	// Out provides either the externally defined true labels or carries the
	// predicted labels.
	Out []float32
}

func (c Config) Verify() {}
