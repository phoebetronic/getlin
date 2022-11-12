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

# implementation

Getlin defines learnable systems as Modules, which in turn are composed of
voting Clauses, which in turn are composed of adaptable Automas.

A Clause learns probability patterns using two polarity sets of Regression
Tsetlin Automata (Automas). See the `automa` package. The two polarities are
positive and negative. The number of Automas per polarity is always equal to the
length of the input Vector `I`, which the Clause receives as input. A Clause
processes input in its original form, here namely positive polarity, and its
reversed form, here namely negative polarity.

In the figure below, Automas track their states along the horizontal state
distribution. Each Automa is only ever responsible for one and the same feature
weight index. Excluded state, between 0.0 and 0.5 along the state distribution,
causes the Automas to exclude the feature weight they are responsible for
respectively. Weights excluded below are denoted with `e`. Conversely, included
state, between 0.5 and 1.0 along the state distribution, causes Automas to
include the feature weights they are responsible for respectively. Weights
included below are denoted with `i`.

The Clause learns by updating the internal state pointers of its Automas, and
predicts by mapping included feature weights to the average of its internal
state pointers, below shown as the output Vector `O`.

Should any negative polarity Automa include a feature weight above 0.5, or
should any positive polarity Automa include a feature weight below 0.5, the
Clause returns 0.

![Clause](/assets/clause.svg)

### classification

One of the most basic pattern recognition tasks is to differentiate classes of
bit patterns. Consider a 4 bit pattern like `[0 1 0 1]`, of which there are 16
unique patterns to differentiate. Getlin trains 1 Module, that is one Tsetlin
Machine per class. The following multi class classification example can be run
and verified via `go test ./ensure/module/multic -v` printing the mean absolute
error for the given epoch. In this very basic example the error is already zero
after the first epoch.

```
epo  1    mae 0.000
epo  2    mae 0.000
epo  3    mae 0.000
epo  4    mae 0.000
epo  5    mae 0.000
epo  6    mae 0.000
epo  7    mae 0.000
epo  8    mae 0.000
epo  9    mae 0.000
epo 10    mae 0.000
epo 11    mae 0.000
epo 12    mae 0.000
epo 13    mae 0.000
epo 14    mae 0.000
epo 15    mae 0.000
```

The final verification message is being printed, showing bit pattern `[0 0 1 0]`
to be class 1 to match for and bit pattern `[1 1 0 1]` to be class 2 to match
for, while classifying all other bit patterns to be none of those. Note the one
hot encoding of the output Vector confirming the class labels respectively.

```
The test data defines [1 0 1 0] in class 0 to be [0], which the Module confirms with [0 0].
The test data defines [0 0 1 0] in class 0 to be [1], which the Module confirms with [1 0].

The test data defines [0 1 1 0] in class 1 to be [0], which the Module confirms with [0 0].
The test data defines [1 1 0 1] in class 1 to be [1], which the Module confirms with [0 1].
```
