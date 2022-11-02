package getlin

type Module interface {
	// Mapper is an internal data structure manager enabling Modules of complex
	// architectures to be updated and rendered.
	Mapper() Mapper
	// Metric returns the metric object carried or constructured by this Module.
	// Tracking of runtime data is specific to each Module implementation.
	Metric() Metric
	// Search is doing inference, computing an output vector given an input
	// vector.
	Search(Vector) Vector
	// Shaper returns the shape specification of this Module.
	Shaper() Shaper
	// States returns the internal state pointers of this Module's TAs as a
	// single list. States may origin from different clauses, which in turn
	// provide their states in different polarities.
	States() []float32
	// Update uses the given vector's true labels for improving all underlying
	// Automa states.
	Update(Vector)
}
