package clause

import "github.com/phoebetron/getlin/metric"

func (c *Clause) Metric() metric.Interface {
	return c.met
}
