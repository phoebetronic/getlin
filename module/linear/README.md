# linear

The Linear Module contains a configured amount of Clauses of which each takes
the binary input Vector `I`.

On the left side of the figure below is shown how search input is transparently
passed to each and every Clause within the Linear Module. Each and every Clause
within the Linear Module produces a single output Vector bit during inference.
All of these single bit inference results are concatenated to form a certain
output Vector according to the Linear Module's output shape. Calling the Linear
Module's `Search` method will transparently collect and return the concatenated
sequence of output Vector `O` as produced by each and every Clause.

On the right side of the figure below is shown how update input is transparently
passed to each and every Clause within the Linear Module. The difference here,
compared to inference as described above, is that the given input Vector must
contain the appropriate true label for each and every Clause, at the correct
true label index matching the index of the Clause to be conditioned. For
instance, while the input Vector bits are all the same for each Clause when
calling the Linear Module's `Update` method, updating Clause number 4 would
require the true label required to condition this very Clause to be positioned
at true label index 3 within the output Vector.

![Linear](/assets/linear.svg)
