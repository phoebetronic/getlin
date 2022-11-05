# ensure

A dedicated package for certain integration test scenarios. While the interface
design of this project is rather clean and underlying implementations are
decoupled sufficiently, certain test cases mess with the dependency tree. In
consequence, to prevent import cycles and other design flaws, more complicated
test cases going beyond simple unit test requirements are put in their dedicated
location.