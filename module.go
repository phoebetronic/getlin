package getlin

type Module interface {
	// Search is doing inference, computing an output Vector given an input
	// Vector.
	Search(Vector) Vector
	// Shaper returns the shape specification of this Module.
	Shaper() Shaper
	// Update uses the given Vector's true weights for refining all underlying
	// Clauses and Automas.
	Update(Vector)
	// Verify returns metrics about the runtime performance given a batch of
	// test Vectors.
	Verify([][2]Vector) Metric
}
