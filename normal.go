package getlin

type Normal interface {
	// Normal is a normalization method computing an output vector given an
	// input vector. Transformation between input and output depends on the
	// underlying implementation. For instance, if the input vector were [0, 1,
	// 0] and Normal were to implement the logical conjunction 0 ∧ 1 ∧ 0, then
	// the output vector would be [0].
	Normal(Vector) Vector
}
