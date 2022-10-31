package getlin

type Vector interface {
	// Cop returns a copy of this vector including all of the same raw bits and
	// true labels.
	Cop() Vector
	// Eql asserts whether the current and the provided vectors are equal. That
	// check verifies the values of carried bits and true labels respectively.
	Eql(Vector) bool
	// Inp returns the binary interface implementation carrying the configured
	// input bits. For any first call or iteration, this might be externally
	// defined binary features. Otherwise the bits carried here are generated
	// output vectors.
	Inp() Binary
	// Out returns the binary interface implementation carrying the configured
	// output bits. For any first call or iteration, this might be externally
	// defined true labels. Otherwise the bits carried here are generated output
	// vectors.
	Out() Binary
}

type Binary interface {
	// Add allows to add a single bit to the current vector.
	Add(bool)
	// And returns the logical conjunction of the current vector by ANDing all
	// of its carried bits. The resulting vector will therefore always only
	// contain a single bit. That is, [1] if all of the current's vector bits
	// are 1s. Otherwise, And returns [0].
	And() Binary
	// Cop returns a copy of this binary implementation including all of the
	// same raw bits.
	Cop() Binary
	// Eql asserts whether the current and the provided binary interfaces carry
	// the same raw vector data for input and output.
	Eql(Binary) bool
	// Neg returns the negative polarity bit carried at the given vector index.
	// If vector were [0, 1, 1, 0] and ind were 2, then Neg would return false.
	Neg(int) bool
	// One returns true in case the current vector contains only 1s and no 0s.
	// If the current vector is empty, then One returns false.
	One() bool
	// Pos returns the negative polarity bit carried at the given vector index.
	// If vector were [0, 1, 1, 0] and ind were 2, then Pos would return true.
	Pos(int) bool
	// Raw returns a copy of all the raw vector bits.
	Raw() []bool
	// Spl cuts the current raw bits at the given index position and returns a
	// new instance for both the left and right side of the split.
	Spl(int) (Binary, Binary)
	// Zer returns true in case the current vector contains only 0s and no 1s.
	// If the current vector is empty, then Zer returns false.
	Zer() bool
}
