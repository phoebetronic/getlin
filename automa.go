package getlin

type Automa interface {
	// Add increases the current state pointer by the given delta, while
	// maintaining the highest value 1.
	Add(float32)
	// Exc reports whether the Automa decided to be excluded based on its
	// current pointer position. That is, any pointer position smaller than 0.
	Exc() bool
	// Hig reports whether the current state pointer reached its highest point
	// already. This can be used to identify if stochastic feedback is even
	// worth pursuing, since randomness comes at a certain cost.
	Hig() bool
	// Inc reports whether the Automa decided to be included based on its
	// current pointer position. That is, any pointer position greater than 0.
	Inc() bool
	// Low reports whether the current state pointer reached its lowest point
	// already. This can be used to identify if stochastic feedback is even
	// worth pursuing, since randomness comes at a certain cost.
	Low() bool
	// Poi returns the current state pointer of an Automa, which will be within
	// the range [0.0, 1.0].
	Poi() float32
	// Red decreases the current state pointer by the given delta, while
	// maintaining the lowest value 0.
	Red(float32)
}
