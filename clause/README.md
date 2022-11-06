# clause

A Clause is formed by two polarity sets of Tsetlin Automata (TA). See the
`automa` package. The two polarities are positive and negative. The number of
TAs is always equal to the length of the binary feature Vector `I`, which the
Clause receives as input. A Clause processes input in its original form, here
namely positive polarity, and its inverted form, here namely negative polarity.

In the figure below, TAs track their states along the horizontal state
distribution. Each TA is only ever responsible for one and the same input bit
index. Excluded state, between 0 and 2 along the state distribution, causes the
TAs to exclude the input bits they are responsible for respectively. Bits
excluded below are denoted with `e`. Conversely, included state, between 3 and 5
along the state distribution, causes TAs to include the input bits they are
responsible for respectively. Bits included below are denoted with `i`. 

The Clause predicts a label via the logical conjunction ^0 ∧ 1 ∧ 0, which
results in the output Vector `O`. Here ^0 is the negated form of 1, both of
which share the same input bit index.

![Clause](/assets/clause.svg)
