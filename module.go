package getlin

type Module interface {
	// Search is doing inference, computing an output Vector given an input
	// Vector.
	Search(Vector) Vector
	// Shaper returns the shape specification of this Module.
	Shaper() Shaper
	// Update uses the given Vector's true labels for improving all underlying
	// Automa states.
	Update(Vector)
	// Verify returns metrics about the runtime performance given a batch of
	// test Vectors. A batch may be a list of 64 sample pairs, of which the
	// first item, index 0, is the negative class and the second item, index 1,
	// is the positive class.
	Verify([][2]Vector) Metric
}
