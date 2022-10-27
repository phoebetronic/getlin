package clause

import (
	"math/rand"

	"github.com/phoebetron/getlin/metric"
	"github.com/phoebetron/getlin/vector"
)

func (c *Clause) Update(vec vector.Interface) {
	var ran float32
	{
		ran = rand.Float32()
	}

	var tru bool
	{
		tru = vec.Tru(0)
	}

	for i := 0; i < vec.Len(); i++ {
		var nrt float32
		var prt float32
		{
			nrt = c.neg[i].Rat()
			prt = c.pos[i].Rat()
		}

		// The negative and positive feedback probability is computed by the
		// progression threshold of the automa, some random distribution and the
		// configured activation function mapping both.
		var nfp bool
		var pfp bool
		{
			nfp = c.act.Active(nrt, ran)
			pfp = c.act.Active(prt, ran)
		}

		if !nfp && !pfp {
			continue
		}

		// The negative and positive label match tells us if the predicted
		// literal label of the automa position index matches the true vector
		// label. If these match respectively, they are desireable. If they
		// don't, we want to correct them probabilistically.
		var nlm bool
		var plm bool
		{
			nlm = vec.Neg(i) == tru
			plm = vec.Pos(i) == tru
		}

		// The negative and positive exclusion and inclusion statements help us
		// to nudge the TAs in more desired directions during the feedback
		// process. We want to include what should be included, therefore what
		// is not included is being voted "up". Conversely, we want to exclude
		// what should be excluded, therefore what is not excluded is being
		// voted "down".
		var nex bool
		var nne bool
		var nin bool
		var pex bool
		var pne bool
		var pin bool
		{
			nex = c.neg[i].Exc()
			nne = c.neg[i].Neu()
			nin = c.neg[i].Inc()
			pex = c.pos[i].Exc()
			pne = c.pos[i].Neu()
			pin = c.pos[i].Inc()
		}

		// Initialization and redistribution occurs if TAs are neutral. This
		// process may apply various methods. For now the default is to set the
		// internal state pointer to a random value along one side of the state
		// distribution. The direction along the state distribution is
		// determined by the desired outcome as provided by the true label.
		if tru && nne {
			c.neg[i].Ini(-1)
		}
		if !tru && nne {
			c.neg[i].Ini(+1)
		}
		if tru && pne {
			c.pos[i].Ini(+1)
		}
		if !tru && pne {
			c.pos[i].Ini(-1)
		}

		// Apply feedback using the information we prepared above. Feedback is
		// applied in case the negative or positive feedback probability
		// triggers respectively. If feedback is triggered, the following
		// conditions and respective actions apply.
		//
		//     If the predicted and true label do not match, while the wrongly
		//     predicted label is included, we discourage inclusion further.
		//
		//     If the predicted and true label do match, while the correctly
		//     predicted label is included, we encourage inclusion further.
		//
		//     If the predicted and true label do match, while the correctly
		//     predicted label is excluded, we encourage inclusion further.
		//
		if !nlm && nin {
			c.neg[i].Rem(1)
		}
		if !plm && pin {
			c.pos[i].Rem(1)
		}
		if nlm && nin {
			c.neg[i].Add(1)
		}
		if plm && pin {
			c.pos[i].Add(1)
		}
		if nlm && nex {
			c.neg[i].Add(1)
		}
		if plm && pex {
			c.pos[i].Add(1)
		}

		// Sample metrics 10% of the time. Most relevant metrics for now are
		// related to computing a confusion matrix and getting insights into the
		// internal automa state distribution.
		if ran < 0.1 {
			if tru == vec.Neg(i) {
				c.met.Set().Mat(metric.TP, 1)
			}
			if !tru == !vec.Neg(i) {
				c.met.Set().Mat(metric.TN, 1)
			}
			if tru == !vec.Neg(i) {
				c.met.Set().Mat(metric.FN, 1)
			}
			if !tru == vec.Neg(i) {
				c.met.Set().Mat(metric.FP, 1)
			}

			if tru == vec.Pos(i) {
				c.met.Set().Mat(metric.TP, 1)
			}
			if !tru == !vec.Pos(i) {
				c.met.Set().Mat(metric.TN, 1)
			}
			if tru == !vec.Pos(i) {
				c.met.Set().Mat(metric.FN, 1)
			}
			if !tru == vec.Pos(i) {
				c.met.Set().Mat(metric.FP, 1)
			}

			if nex {
				c.met.Set().Sta(-nrt)
			}
			if nrt == 0 {
				c.met.Set().Sta(nrt)
			}
			if nin {
				c.met.Set().Sta(+nrt)
			}

			if pex {
				c.met.Set().Sta(-prt)
			}
			if prt == 0 {
				c.met.Set().Sta(prt)
			}
			if pin {
				c.met.Set().Sta(+prt)
			}
		}
	}
}
