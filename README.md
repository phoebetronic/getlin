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

### implementation

The first implementation of the Tsetlin Automata states here is based on Golang
`bool` types, of which each occupies a single byte. See the `clause` and
`automa` packages. In fact, the smallest addressable unit in Golang is a single
byte, with the exception of bitwise operations on such a single byte. There is
no way I know about to allocate and manage single bits conveniently in Golang as
is. Therefore the current implementation settled with lists of booleans.

In theory, states of TA Teams may be managed in groups of 8 bits as described in
the original paper. Single bit state can be represented and manipulated
individually at will, getting most bang for your buck during memory allocation
and state storage. This merely comes at the cost of code design and
comprehension. For now I consider this idea to be too complicated to worry about
just starting out. Contributions are gladly welcome.
