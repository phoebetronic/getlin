package getlin

type Shaper interface {
	// Inp is the input vector length specified to be accepted.
	Inp() int
	// Out is the output vector length specified to be accepted.
	Out() int
	// Sta is the amount of internal state pointers managed.
	Sta() int
}
