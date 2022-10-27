package clause

import (
	"github.com/phoebetron/getlin/active"
	"github.com/phoebetron/getlin/metric"
	"github.com/phoebetron/getlin/serial"
)

type Config struct {
	// Act is an activation function used for stochastic feedback activation.
	Act active.Interface
	// Met is the telemetric like object recording statistical information about
	// the Clause's runtime behaviour.
	Met metric.Interface
	// Ser is a serialization method enabling the Clause to be stored and
	// loaded. Algorithms for JSON or Protocl Buffers may be used here.
	Ser serial.Interface
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
	// is the number of binary features of the desired input vectors to train.
	// That is, input vectors represented with 4 bits require 4 TAs, e.g. 0110.
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

func (c *Config) Verify() {
	if c.Act == nil {
		panic("Config.Act must not be empty")
	}
	if c.Met == nil {
		panic("Config.Met must not be empty")
	}
	// if c.Ser == nil {
	// 	panic("Config.Ser must not be empty")
	// }
	if c.Sta == 0 {
		panic("Config.Sta must not be empty")
	}
	if c.Tas == 0 {
		panic("Config.Tas must not be empty")
	}
}
