package clause

import "github.com/phoebetron/getlin"

type Config struct {
	// Fre is the probability used to calculate the feedback frequency in two
	// different ways, namely high and low. Consider the following rules for Fre
	// being set to 0.1.
	//
	//    The low frequency updates happen with a probability of 0.1.
	//
	//    The high frequency updates happen with a probability of 1 - 0.1.
	//
	Fre float32
	// Ran provides randomization primitives for applying stochastic feedback.
	Ran getlin.Random
	// Sta is the number of states S along a single side of the states
	// distribution, where the final states distribution will be (2 * S) + 1.
	// Below is illustrated a state distribution 9 = S = 4.
	//
	//                    -                                     +
	//
	//           S     3     2     1         0         1     2     3     S
	//
	//           ·     ·     ·     ·         ·         ·     i     ·     ·
	//
	//           ·     e     ·     ·         ·         ·     ·     ·     ·
	//
	//           ·     ·     ·     ·         n         ·     ·     ·     ·
	//
	//           ·     ·     ·     e         ·         ·     ·     ·     ·
	//
	Sta int
	// Tas is the number of Tsetlin Automata (TAs) managed by the Clause. This
	// is the number of binary features of the desired input Vectors to train.
	// That is, input Vectors represented with 4 bits require 4 TAs, e.g. 0110.
	// In the standard implementation the total amount of TAs managed by the
	// Clause is 2 * Tas. That is, one set for the positive polarity and one set
	// for the negative polarity. Below is illustrated a TA Array 8 = Tas = 4.
	//
	//           1     ·     ·     ·         ·         ·     i     ·     ·
	//
	//     -     0     e     ·     ·         ·         ·     ·     ·     ·
	//
	//           0     ·     ·     ·         n         ·     ·     ·     ·
	//
	//           1     ·     ·     e         ·         ·     ·     ·     ·
	//
	//
	//           0     ·     ·     ·         ·         ·     ·     i     ·
	//
	//     +     1     ·     e     ·         ·         ·     ·     ·     ·
	//
	//           1     ·     ·     ·         ·         ·     i     ·     ·
	//
	//           0     ·     ·     ·         n         ·     ·     ·     ·
	//
	Tas int
}

func (c Config) Verify() {
	if c.Fre == 0 {
		panic("Config.Fre must not be empty")
	}
	if c.Ran == nil {
		panic("Config.Ran must not be empty")
	}
	if c.Sta == 0 {
		panic("Config.Sta must not be empty")
	}
	if c.Tas == 0 {
		panic("Config.Tas must not be empty")
	}
}
