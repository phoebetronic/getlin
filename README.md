# getlin

A Golang implementation of the [Tsetlin Machine] as proposed in [the original
paper]. So `G`olang + Ts`etlin` = `Getlin`, as in [Gatling Gun], you get it.

### abstract

Given a binary input vector, the system produces a certain output vector. The
output vector quality improves upon learning predictive properties via
reinforcement strategies. The system captures patterns within training data
based on encapsulated results of [Boolean Algebra] without the need of complex
matrix multiplications. These results are combined and voted upon, concluding in
learnable global maxima.

[Tsetlin Machine]: https://en.wikipedia.org/wiki/Tsetlin_machine
[the original paper]: https://arxiv.org/pdf/1804.01508.pdf
[Gatling Gun]: https://en.wikipedia.org/wiki/Gatling_gun
[Boolean Algebra]: https://en.wikipedia.org/wiki/Boolean_algebra

### classification

One of the most basic pattern recognition tasks is to differentiate classes of
bit patterns. Consider a 4 bit pattern like `[0 1 0 1]`, of which there are 16
unique patterns to differentiate. Getlin trains 1 Module, that is one Tsetlin
Machine per class. The following multi class classification example can be run
and verified via `go test ./ensure/module/multic -v` printing the mean absolute
error and the accuracy for the given epoch.

```
epo  1    mae 0.117    acc 0.883
epo  2    mae 0.000    acc 1.000
epo  3    mae 0.000    acc 1.000
epo  4    mae 0.000    acc 1.000
epo  5    mae 0.000    acc 1.000
```

The final verification message is being printed, showing bit pattern `[0 0 1 0]`
to be class 1 to match for and bit pattern `[1 1 0 1]` to be class 2 to match
for, while classifying all other bit patterns to be none of those. Note the one
hot encoding of the output Vector confirming the class labels respectively.

```
The test data defines [1 0 1 0] in class 1 to be [0], which the Module confirms with [0 0].
The test data defines [0 0 1 0] in class 1 to be [1], which the Module confirms with [1 0].

The test data defines [0 1 1 0] in class 2 to be [0], which the Module confirms with [0 0].
The test data defines [1 1 0 1] in class 2 to be [1], which the Module confirms with [0 1].
```

### implementation

The first implementation of the Tsetlin Automata states here is based on Golang
`uint8` types, of which each occupies a single byte. See the `clause` and
`automa` packages. In fact, the smallest addressable unit in Golang is a single
byte, with the exception of bitwise operations on such a single byte. There is
no way I know about to allocate and manage single bits conveniently in Golang as
is. Therefore the current implementation settled with lists of integers.

In theory, states of TA Teams may be managed in groups of 8 bits as described in
the original paper. Single bit state can be represented and manipulated
individually at will, getting most bang for your buck during memory allocation
and state storage. This merely comes at the cost of code design and
comprehension and maybe even execution time. For now I consider this idea to be
too complicated to worry about just starting out. Contributions though are
always gladly welcome.
