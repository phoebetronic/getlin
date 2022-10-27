# clause 

A Clause is formed by collections of Tsetlin Automata (TA). See the `automa`
package. The number of TAs is always equal to the length of the binary input
vector `X`, or the number of features processed by each and every Clause. Each
Clause processes input in its original form, here namely positive polarity, and
its inverted form, here namely negative polarity. 

In the figure below, TAs track their states along the horizontal state
distribution. Each TA is only ever responsible for one and the same input bit
index. Negative state, left of 0 along the state distribution, causes the TAs to
exclude the input bits they are responsible for respective. Bits excluded below
are denoted with `e`. Conversely, positive state, right of 0 along the state
distribution, causes TAs to include the input bits they are responsible for
respectively. Bits included below are denoted with `i`. The Getlin
implementation here diverges from the original paper by adding the neutral state
`n`, which is the TA state exactly at 0. This additional differentiation is
practical during feedback activation.

The Clause's output vector is calculated by normalizing all included bits by
means of the normalization function `N`, which, as shown below, is the logical
conjunction 0 ∧ 1 ∧ 0. As per the Getlin implementation, in case any TA decides
to include a 0, the whole Clause returns 0.

![Clause](/assets/clause.svg)