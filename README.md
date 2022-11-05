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
bit patterns. Consider a 4 bit pattern like `[1 1 0 1]`, of which there are 16
unique patterns to differentiate. Getlin trains 1 Module, that is one Tsetlin
Machine per class. The following examples can be run and verified via `go test
./ensure/module/voting -v`.

Running the first test prints the mean absolute error and the accuracy for the
given epoch.

```
epo  1    mae 0.333    acc 0.667
epo  2    mae 0.133    acc 0.867
epo  3    mae 0.000    acc 1.000
epo  4    mae 0.000    acc 1.000
epo  5    mae 0.000    acc 1.000
epo  6    mae 0.000    acc 1.000
epo  7    mae 0.000    acc 1.000
epo  8    mae 0.000    acc 1.000
epo  9    mae 0.000    acc 1.000
epo 10    mae 0.000    acc 1.000
```

The final verification message is being printed, showing bit pattern `[1 1 0 1]`
to be the class to match for, while classifying all other bit patterns to be
false.

```
The test data defines [0 1 1 0] to be 0, which the Module confirms with 0.
The test data defines [1 1 0 1] to be 1, which the Module confirms with 1.
```

Running the second test prints the mean absolute error and the accuracy for the
given epoch. We see a slightly different progression until the Module matches
perfectly.

```
epo  1    mae 0.267    acc 0.733
epo  2    mae 0.233    acc 0.767
epo  3    mae 0.233    acc 0.767
epo  4    mae 0.167    acc 0.833
epo  5    mae 0.067    acc 0.933
epo  6    mae 0.033    acc 0.967
epo  7    mae 0.067    acc 0.933
epo  8    mae 0.000    acc 1.000
epo  9    mae 0.000    acc 1.000
epo 10    mae 0.000    acc 1.000
```

The final verification message is being printed, showing bit pattern `[0 0 1 0]`
to be the class to match for, while classifying all other bit patterns to be
false.

```
The test data defines [1 0 1 0] to be 0, which the Module confirms with 0.
The test data defines [0 0 1 0] to be 1, which the Module confirms with 1.
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
