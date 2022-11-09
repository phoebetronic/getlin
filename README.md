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
and verified via `go test ./ensure/module/multic -v` printing the converging
mean absolute error and the accuracy for the given epoch.

```
epo  1    mae 0.200    acc 0.800
epo  2    mae 0.133    acc 0.867
epo  3    mae 0.100    acc 0.900
epo  4    mae 0.067    acc 0.933
epo  5    mae 0.000    acc 1.000
epo  6    mae 0.017    acc 0.983
epo  7    mae 0.017    acc 0.983
epo  8    mae 0.017    acc 0.983
epo  9    mae 0.000    acc 1.000
epo 10    mae 0.000    acc 1.000
epo 11    mae 0.000    acc 1.000
epo 12    mae 0.000    acc 1.000
epo 13    mae 0.000    acc 1.000
epo 14    mae 0.000    acc 1.000
epo 15    mae 0.000    acc 1.000
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
