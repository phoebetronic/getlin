package getlin

type Random interface {
	// F32 returns a number in the half-open interval [0.0,1.0) generated from
	// the underlying source, which may or may not be truly random.
	F32() float32
}
