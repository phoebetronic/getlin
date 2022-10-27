package automa

type Interface interface {
	// Add increases the automa's state distribution by cou. That is, if the
	// automa's internal state pointer were 15 and cou were 3, then the automa's
	// internal state pointer would be set to 18. Similarly adding 5 to the
	// pointer -17 would result in the slightly higher pointer -12.
	Add(float32)
	// Exc reports whether the automa decides to be excluded based on its
	// current pointer position. That is, any pointer position smaller than 0.
	Exc() bool
	// Inc reports whether the automa decides to be included based on its
	// current pointer position. That is, any pointer position greater than 0.
	Inc() bool
	// Ini sets the internal state pointer to a randomly selected value along
	// the automa's state distribution, in direction of the given indicator,
	// either +1 for a positive state pointer, or -1 for a negative state
	// pointer.
	Ini(int)
	// Neu reports whether the automa decides to be neutral based on its current
	// pointer position. That is, the special pointer position exactly at 0.
	Neu() bool
	// Poi returns the automa's internal state as currently configured.
	Poi() float32
	// Rat returns the progression of the target ratio between the current
	// pointer position the the absolute state boundary. This ratio may be used
	// for stochastic feedback activation. For instance, if the lower state
	// boundary were -25 and the current pointer position were -23, then Rat
	// would return poi / sta = 0.92. That means that the automa has already
	// progressed 92% into the direction of the most outer boundary possible.
	Rat() float32
	// Rem decreases the automa's state distribution by cou. That is, if the
	// automa's internal state pointer were 15 and cou were 3, then the automa's
	// internal state pointer would be set to 12. Similarly removing 5 from the
	// pointer -17 would result in the even lower pointer -22.
	Rem(float32)
}
