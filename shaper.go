package getlin

type Shaper interface {
	// Cla is the amount of Clauses to be utilized by a Module.
	Cla() int
	// Inp is the input vector length specified to be received by a Module.
	Inp() int
	// Out is the output vector length specified to be returned by a Module.
	Out() int
	// Sta is the amount of internal state pointers managed by a Module.
	Sta() int
}
