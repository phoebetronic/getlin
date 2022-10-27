package active

type hrl struct{}

// Hrl is a form of hard probabilistic ReLu activation. The minimum threshold
// may not be lower than 0.05 in any case. That ensures for instance neutral
// automas having their internal pointer states exactly at zero to be updated at
// least 5% of the time. Above that minimum threshold, activation probability
// grows linearly until it peaks at 1.0 like a classical ReLu.
func Hrl() Interface {
	return &hrl{}
}

func (q *hrl) Active(rat float32, ran float32) bool {
	if rat <= 0.05 {
		return ran <= 0.05
	}

	return ran <= rat
}
