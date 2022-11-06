package getlin

type Clause interface {
	// Output produces the Clause's binary label prediction given a binary
	// feature Vector.
	Output([]uint8) uint8
	// States returns the internal state pointers of this Clause's TAs. The
	// first list returned represents states of negative polarity and the second
	// list returned represents states of positive polarity.
	States() ([]int, []int)
	// TypOne applies Type 1 Feedback to the Clause, given a binary feature
	// Vector and its own label prediction.
	TypOne([]uint8, uint8)
	// TypTwo applies Type 2 Feedback to the Clause, given a binary feature
	// Vector and its own label prediction.
	TypTwo([]uint8, uint8)
}
