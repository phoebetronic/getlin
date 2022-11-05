package harrel

import "github.com/phoebetron/getlin"

// Harrel is a form of hard ReLu activation. The minimum threshold may not be
// lower than 0.05 in any case. That ensures for instance neutral Automas having
// their internal pointer states exactly at zero to be updated at least 5% of
// the time. Above that minimum threshold, activation probability grows linearly
// like a classical ReLu.
type Harrel struct{}

func New(con Config) getlin.Active {
	{
		con.Verify()
	}

	return &Harrel{}
}

func (q *Harrel) Act(rat float32, prb float32) bool {
	if rat <= 0.05 {
		return prb <= 0.05
	}

	return prb <= rat
}
