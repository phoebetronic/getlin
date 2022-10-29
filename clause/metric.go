package clause

import "github.com/phoebetron/getlin"

func (c *Clause) Metric() getlin.Metric {
	return c.met
}
