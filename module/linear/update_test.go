package linear

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/random/static"
	"github.com/phoebetron/getlin/vector"
)

func Test_Module_Linear_Update(t *testing.T) {
	testCases := []struct {
		inp []int
		tru []int
		ini []float32
		upd []float32
	}{
		// Case 0
		{
			inp: []int{
				1,
				1,
			},
			tru: []int{
				0,
				0,
				0,
			},
			ini: []float32{
				+2, +4, -4, -4,
				+2, +2, -3, -1,
				+1, +3, -1, -4,
			},
			// Updating the Linear Module increased the states of negative
			// polarity, because the negated form of input [1 1] is supposed to
			// include [0 0] in each and every Clause based on the true labels
			// [0 0 0]. In other words, all the states of negative polarity
			// shift towards more inclusion because the true labels are all
			// negative.
			//
			//     [-   -] [0] [+   +]
			//      +
			//      +   +
			//      +   +
			//
			upd: []float32{
				+3, +4, -4, -4,
				+3, +3, -3, -1,
				+2, +4, -1, -4,
			},
		},
		// Case 1
		{
			inp: []int{
				1,
				1,
			},
			tru: []int{
				1,
				0,
				0,
			},
			ini: []float32{
				-3, -1, +2, +3,
				+4, +2, -2, -3,
				+4, +1, -3, -3,
			},
			// Updating the Linear Module increased the states of negative
			// polarity for the second and third Clause, because the negated
			// form of input [1 1] is supposed to include [0 0] in those Clauses
			// based on the true labels [1 0 0]. Conversely, Updating the Linear
			// Module increased the states of positive polarity for the first
			// Clause, because the normal form of input [1 1] is supposed to
			// include [1 1] in this Clause based on the true labels [1 0 0]. In
			// other words, all the states of negative polarity shift towards
			// more inclusion where the true labels are false, while all the
			// states of positive polarity shift towards more inclusion where
			// the true labels are true.
			//
			//     [-   -] [0] [+   +]
			//                  +   +
			//          +
			//          +
			//
			upd: []float32{
				-3, -1, +3, +4,
				+4, +3, -2, -3,
				+4, +2, -3, -3,
			},
		},
		// Case 2
		{
			inp: []int{
				0,
				0,
			},
			tru: []int{
				0,
				0,
				0,
			},
			ini: []float32{
				+4, +4, -4, -1,
				+3, +2, -4, -1,
				+4, +2, -4, -1,
			},
			// Updating the Linear Module decreased the states of negative
			// polarity, because the negated form of input [0 0] is supposed to
			// exclude [1 1] in each and every Clause based on the true labels
			// [0 0 0]. In other words, all the states of negative polarity
			// shift towards more exclusion because the true labels are all
			// negative, while all the states of positive polarity shift towards
			// more inclusion where the false labels are false as well.
			//
			//     [-   -] [0] [+   +]
			//      -   -       +   +
			//      -   -       +   +
			//      -   -       +   +
			//
			upd: []float32{
				+3, +3, -3, 0,
				+2, +1, -3, 0,
				+3, +1, -3, 0,
			},
		},
		// Case 3
		{
			inp: []int{
				0,
				0,
			},
			tru: []int{
				0,
				1,
				0,
			},
			ini: []float32{
				+2, +2, -4, -3,
				-2, -3, +3, +4,
				+2, +3, -4, -1,
			},
			// Updating the Linear Module decreased the states of negative
			// polarity in Clause 1 and 3, because the negated form of input [0
			// 0] is supposed to exclude [1 1] in those Clauses based on the
			// true labels [0 1 0], while updating the Linear Module increased
			// the states of negative polarity in Clause 2, because the negated
			// form of input [0 0] is supposed to include [1 1] in this Clauses
			// based on the true labels [0 1 0]. The opposite applies to the
			// states of positive polarity.
			//
			//     [-   -] [0] [+   +]
			//      -   -       +   +
			//      +   +       -   -
			//      -   -       +   +
			//
			upd: []float32{
				+1, +1, -3, -2,
				-1, -2, +2, +3,
				+1, +2, -3, 0,
			},
		},
		// Case 4
		{
			inp: []int{
				1,
				0,
			},
			tru: []int{
				0,
				0,
				1,
			},
			ini: []float32{
				+4, +2, -2, -2,
				+2, +1, -4, -2,
				-1, -4, +3, +4,
			},
			//
			//     [-   -] [0] [+   +]
			//
			//          -           +
			//      +   -           +
			//          +       +   +
			//
			upd: []float32{
				+4, +1, -2, -1,
				+3, 0, -4, -1,
				-1, -3, +4, +3,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var mod getlin.Module
			{
				mod = muslin(2, 3)
			}

			// The first update call does not modify the TA states based on the
			// given true label feedback, because after Module creation all
			// states are neutral and therefore only get randomly initialized.
			{
				mod.Update(musvec(tc.inp, tc.tru))
			}

			var ini []float32
			{
				ini = mod.States()
			}

			// After initialization, the second update call modifies the TA
			// states on the given true label feedback deterministically, at
			// stochastically chosen points in time.
			{
				mod.Update(musvec(tc.inp, tc.tru))
			}

			var upd []float32
			{
				upd = mod.States()
			}

			if !reflect.DeepEqual(ini, tc.ini) {
				t.Fatalf("ini\n\n%s\n", cmp.Diff(tc.ini, ini))
			}
			if !reflect.DeepEqual(upd, tc.upd) {
				t.Fatalf("upd\n\n%s\n", cmp.Diff(tc.upd, upd))
			}
		})
	}
}

func muslin(inp int, out int) getlin.Module {
	return New(Config{
		Inp: inp,
		Out: out,
		Ran: static.New(static.Config{F32: 0.05}),
		Sta: 4,
	})
}

func musvec(inp []int, out []int) getlin.Vector {
	return vector.New(vector.Config{
		Inp: vector.ToBool(inp...),
		Out: vector.ToBool(out...),
	})
}
