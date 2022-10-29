package getlin

type Vector interface {
	// Add allows to add a single bit to the current vector.
	Add(bit bool)
	// And returns the logical conjunction of the current vector by ANDing all
	// of its carried bits. The resulting vector will therefore always only
	// contain a single bit. That is, [1] if all of the current's vector bits
	// are 1s. Otherwise, And returns [0].
	And() Vector
	// Bit returns a copy of the raw vector inputs.
	Bit() []bool
	// Eql asserts whether the current and the provided vectors are equal. That
	// check verifies the values of carried bits and true labels respectively.
	Eql(Vector) bool
	// Neg returns the negative polarity bit carried at the given vector index.
	// If vector were [0, 1, 1, 0] and ind were 2, then Neg would return false.
	Neg(int) bool
	// One returns true in case the current vector contains only 1s and no 0s.
	// If the current vector is empty, then One returns false.
	One() bool
	// Pos returns the negative polarity bit carried at the given vector index.
	// If vector were [0, 1, 1, 0] and ind were 2, then Pos would return true.
	Pos(int) bool
	// Tru returns a copy of the true vector labels.
	Tru() []bool
	// Zer returns true in case the current vector contains only 0s and no 1s.
	// If the current vector is empty, then Zer returns false.
	Zer() bool
}
