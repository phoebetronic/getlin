# clause

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
