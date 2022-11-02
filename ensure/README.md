# ensure

A dedicated package for certain integration test scenarios. While the interface
design of this project is rather clean and underlying implementations are
decoupled sufficiently, certain test cases mess with the dependency tree. For
instance the `mapper` tests require the `graphs` implementation in order to
ensure the correct mapping mechanisms to work as expected. In turn, the `graphs`
tests require the `mapper` implementation in order to ensure the correct search
and update mechanisms to work as expected. In consequence, to prevent import
cycles and other design flaws, more complicated test cases going beyond simple
unit test requirements are put in their dedicated location.