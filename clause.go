package getlin

type Clause interface {
	// Output produces the Clause's binary label prediction given a binary
	// feature Vector.
	Output([]uint8) uint8
	//
	States() ([]int, []int)
	// TypOne applies Type 1 Feedback to the Clause, given a binary feature
	// Vector and its own label prediction.
	TypOne([]uint8, uint8)
	// TypTwo applies Type 2 Feedback to the Clause, given a binary feature
	// Vector and its own label prediction.
	TypTwo([]uint8, uint8)
}
