package vector

type Config struct {
	// Cla provides optional class identifier information for different Module
	// implementations to leverage internally.
	Cla int
	// Inp carries the input bits consumed by any given Module, be it inference
	// or training.
	Inp []uint8
	// Out provides either the externally defined true labels or carries the
	// predicted labels.
	Out []uint8
}

func (c Config) Verify() {}
