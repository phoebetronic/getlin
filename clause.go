package getlin

type Clause interface {
	// Output produces the Clause's weight prediction given probabilistic
	// feature Vector. Features and prediction are within the range [0.0, 1.0].
	Output([]float32) float32
	// States returns the current state pointers of this Clause's Automas within
	// the range [0.0, 1.0]. The first list returned represents states of
	// negative polarity and the second list returned represents states of
	// positive polarity.
	States() ([]float32, []float32)
	// TypOne applies Type 1 Feedback to the Clause in order to drive prediction
	// weights towards 1, so that patterns are included and selected for. TypOne
	// takes probabilistic feature weights and its own weight prediction as
	// input. Features and prediction are within the range [0.0, 1.0].
	TypOne([]float32, float32)
	// TypTwo applies Type 2 Feedback to the Clause in order to drive prediction
	// weights towards 0, so that patterns are excluded and discriminated
	// against. TypTwo takes probabilistic feature weights and its own weight
	// prediction as input. Features and prediction are within the range [0.0,
	// 1.0].
	TypTwo([]float32, float32)
}
