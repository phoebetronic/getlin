package getlin

type Vector interface {
	// Cop returns a copy of this Vector including all of the same raw bits and
	// true labels.
	Cop() Vector
	// Eql asserts whether the current and the provided Vectors are equal. That
	// check verifies the values of carried bits and true labels respectively.
	Eql(Vector) bool
	// Ind returns the optional instruction of Module update paths at the root
	// of a Graphs Module. See vector.Config for more information.
	Ind() [][]bool
	// Inp returns the binary interface implementation carrying the configured
	// input bits. For any first call or iteration, this might be externally
	// defined binary features. Otherwise the bits carried here are generated
	// output Vectors.
	Inp() Binary
	// Out returns the binary interface implementation carrying the configured
	// output bits. For any first call or iteration, this might be externally
	// defined true labels. Otherwise the bits carried here are generated output
	// Vectors.
	Out() Binary
	// Sta returns status information of the current Vector. The Status
	// primitives should not be touched by outside users.
	Sta() Status
}

type Binary interface {
	// Add allows to add a single bit to the current Vector.
	Add(uint8)
	// And returns the logical conjunction of the current Vector by ANDing all
	// of its carried bits. The resulting Vector will therefore always only
	// contain a single bit. That is, [1] if all of the current Vector's bits
	// are 1s. Otherwise, And returns [0].
	And() Binary
	// Cop returns a copy of this binary implementation including all of the
	// same raw bits.
	Cop() Binary
	// Cou returns the amount of bits currently held within the Binary buffer
	// matching the given polarity. If Vector were [1 0 1 1 0], then Cou(true)
	// would return 3.
	Cou(uint8) int
	// Eql asserts whether the current and the provided binary interfaces carry
	// the same raw Vector data for input and output.
	Eql(Binary) bool
	// Maj returns the majority bit. If Vector were [1 0 1 1 0], then Maj would
	// return 1. If the amount of 0s and 1s were to be equal, Maj would default
	// to returning 0.
	Maj() uint8
	// Neg returns the negative polarity bit carried at the given Vector index.
	// If Vector were [0 1 1 0] and ind were 2, then Neg would return 0.
	Neg(int) uint8
	// One returns true in case the current Vector contains only 1s and no 0s.
	// If the current Vector is empty, then One returns false.
	One() bool
	// Pos returns the positive polarity bit carried at the given Vector index.
	// If Vector were [0 1 1 0] and ind were 2, then Pos would return 1.
	Pos(int) uint8
	// Raw returns a copy of all the raw Vector bits.
	Raw() []uint8
	// Wei returns the ratio of the given bit across the whole Binary buffer.
	Wei(uint8) float32
	// Zer returns true in case the current Vector contains only 0s and no 1s.
	// If the current Vector is empty, then Zer returns false.
	Zer() bool
}

type Status interface {
	// Fai expresses the current Vector to carry a failed match between true
	// labels and predicted outputs.
	Fai() bool
	// Ini returns true in case the current Vector's status got initialized
	// internally.
	Ini() bool
	// Suc expresses the current Vector to carry a successful match between true
	// labels and predicted outputs.
	Suc() bool
	// Upd modifies the Vector's status based on the provided value. Calling
	// Status.Upd(true) will cause Status.Suc to return true as well. The
	// opposite does then apply to Status.Fai.
	Upd(sta bool)
}
