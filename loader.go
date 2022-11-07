package getlin

type Loader interface {
	// Create executes any setup required to provide batches of training data.
	Create()
	// Search returns a batch of training data in the form of Vector pairs. A
	// batch may be a list of 64 sample pairs, of which the first item, index 0,
	// is the negative class and the second item, index 1, is the positive
	// class.
	Search() [][2]Vector
}
