package getlin

type Mapper interface {
	// Abo returns the list of modules within the previous layer relative to the
	// given module. Consider Module M1 and Module M2 in the first layer and
	// Module M3 in the second layer. Calling Abo(M3) returns M1 and M2.
	//
	//     [
	//         [ M1  M2 ]
	//         [   M3   ]
	//     ]
	//
	Abo(Module) []Module
	// All returns the hierarchical list of modules as configured in this Mapper.
	//
	//     [
	//         [ M1  M2 ]
	//         [   M3   ]
	//     ]
	//
	All() [][]Module
	// Bel returns the list of modules within the next layer relative to the
	// given module. Consider Module M1 and Module M2 in the first layer and
	// Module M3 in the second layer. Calling Bel(M2) returns M3.
	//
	//     [
	//         [ M1  M2 ]
	//         [   M3   ]
	//     ]
	//
	Bel(Module) []Module
	// Ind returns the index of the given module within the context of this
	// Mapper's full graph. Consider Module M1 and Module M2 in the first layer
	// and Module M3 in the second layer. Calling Ind(M2) returns 1.
	//
	//     [
	//         [ M1  M2 ]
	//         [   M3   ]
	//     ]
	//
	Ind(Module) int
	// Lay returns the index of the layer which the given module is part of.
	// Consider Module M1 and Module M2 in the first layer and Module M3 in the
	// second layer. Calling Lay(M2) returns 0.
	//
	//     [
	//         [ M1  M2 ]
	//         [   M3   ]
	//     ]
	//
	Lay(Module) int
	// Tru returns the true label index range as two integer boundaries for the
	// given learnable module, e.g. [3 7]. Learnable modules are updatable,
	// which means they contain clauses and automas to be tuned. Modules not to
	// be updated are purely functional and do not contain any clauses nor
	// automas. For those purely functional modules Tru returns nil. Consider
	// Module M1 and Module M2 in the first layer and Module M3 in the second
	// layer. M1 and M2 have two inputs and one output each, which feed into M3
	// respectively. Calling Tru(M2) returns [1 2], so that the true labels for
	// M2 can be resolved via T[1:2].
	//
	//     I   [      0  1  1  1      ]
	//
	//     [
	//         [ M1(I2,O1)  M2(I2,O1) ]
	//         [      M3(I2,O2)       ]
	//     ]
	//
	//     T   [         0  1         ]
	//
	Tru(Module) []int
}
