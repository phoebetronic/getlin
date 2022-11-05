package getlin

type Automa interface {
	// Add increases the automa's state distribution by 1. That is, if the
	// automa's internal state pointer were 15, then the automa's internal state
	// pointer would be set to 16.
	Add()
	// Exc reports whether the automa decides to be excluded based on its
	// current pointer position. That is, any pointer position smaller than 0.
	Exc() bool
	// Hig reports whether the internal state pointer reached its highest point
	// already. This can be used to identify if stochastic feedback is even
	// worth pursuing, since randomness comes at a certain cost.
	Hig() bool
	// Inc reports whether the automa decides to be included based on its
	// current pointer position. That is, any pointer position greater than 0.
	Inc() bool
	// Low reports whether the internal state pointer reached its lowest point
	// already. This can be used to identify if stochastic feedback is even
	// worth pursuing, since randomness comes at a certain cost.
	Low() bool
	//
	Poi() int
	// Red decreases the automa's state distribution by 1. That is, if the
	// automa's internal state pointer were 15, then the automa's internal state
	// pointer would be set to 14.
	Red()
}
