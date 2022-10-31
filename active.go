package getlin

type Active interface {
	// Act takes the ratio of automa states distributions as first argument, and
	// some ratio of probabilistic random distributions as second argument. The
	// given ratio and probability are used to produce an activation signal for
	// triggering TA feedback.
	Act(float32, float32) bool
}
