package getlin

type Vector interface {
	// Cla returns the optional class identifier this Vector is associated with.
	Cla() int
	// Inp returns the configured feature weights of this Vector.
	Inp() []float32
	// Out returns the configured true weights or weight predictions of this
	// Vector.
	Out() []float32
}
