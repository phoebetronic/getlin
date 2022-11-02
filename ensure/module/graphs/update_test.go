package graphs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Module_Graphs_Update_Metric(t *testing.T) {
	testCases := []struct {
		inp getlin.Vector
		mod getlin.Module
		bef [41]float32
		aft [41]float32
	}{
		// Case 0
		{
			inp: musvec([]int{0, 1}, []int{0, 1}),
			mod: musgra([][]getlin.Module{
				{
					muslin(2, 1),
					muslin(2, 1),
				},
				{
					muslin(2, 2),
				},
			}),
			bef: [41]float32{},
			aft: [41]float32{
				0.1875,
				0,
				0,
				0,
				0,
				0.0625,
				0,
				0,
				0,
				0,
				0.1875,
				0,
				0,
				0,
				0,
				0.0625,
				0,
				0,
				0,
				0, // -
				0, // 0
				0, // +
				0,
				0,
				0,
				0.1875,
				0,
				0,
				0,
				0,
				0.0625,
				0,
				0,
				0,
				0,
				0.1875,
				0,
				0,
				0,
				0,
				0.0625,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var bef [41]float32
			{
				bef = tc.mod.Metric().Get().Sta().Nrm()
			}

			if !reflect.DeepEqual(bef, tc.bef) {
				t.Fatalf("bef\n\n%s\n", cmp.Diff(tc.bef, bef))
			}

			{
				tc.mod.Update(tc.inp)
			}

			var aft [41]float32
			{
				aft = tc.mod.Metric().Get().Sta().Nrm()
			}

			if !reflect.DeepEqual(aft, tc.aft) {
				t.Fatalf("aft\n\n%s\n", cmp.Diff(tc.aft, aft))
			}
		})
	}
}

func Test_Module_Graphs_Update_States(t *testing.T) {
	testCases := []struct {
		inp getlin.Vector
		mod getlin.Module
		bef []float32
		aft []float32
	}{
		// Case 0
		{
			inp: musvec([]int{0, 1}, []int{0, 1}),
			mod: musgra([][]getlin.Module{
				{
					muslin(2, 1),
					muslin(2, 1),
				},
				{
					muslin(2, 2),
				},
			}),
			bef: []float32{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			aft: []float32{
				+4,
				+2,
				-2,
				-3,
				-4,
				-1,
				+3,
				+3,
				+4,
				+4,
				-4,
				-1,
				+3,
				+2,
				-4,
				-1,
			},
		},
		// Case 1
		{
			inp: musvec([]int{0, 1, 0}, []int{0, 1}),
			mod: musgra([][]getlin.Module{
				{
					muslin(3, 2),
					muslin(3, 2),
					muslin(3, 2),
					muslin(3, 2),
				},
				{
					musgra([][]getlin.Module{
						{
							muslin(8, 2),
							muslin(8, 2),
						},
						{
							musvot(4),
						},
					}),
					musgra([][]getlin.Module{
						{
							muslin(8, 2),
							muslin(8, 2),
						},
						{
							musvot(4),
						},
					}),
				},
			}),
			bef: []float32{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			aft: []float32{
				+4, +2, +2, -4, -1, -4, +2, +2,
				+3, -3, -3, -4, +2, +3, +4, -4,
				-1, -2, +2, +2, +1, -2, -4, -2,
				-1, -4, -4, +3, +4, +3, -2, -3,
				-2, +1, +1, +3, -4, -4, -3, +3,
				+1, +3, -4, -2, -1, +3, +4, +1,
				+4, +2, +2, +2, +4, +3, +3, +1,
				-2, -2, -4, -4, -3, -1, -4, -4,
				+1, +2, +2, +4, +4, +3, +3, +3,
				-4, -3, -4, -2, -3, -2, -1, -2,
				+4, +1, +4, +4, +3, +2, +3, +2,
				-3, -2, -2, -3, -3, -4, -3, -3,
				+4, +1, +2, +4, +1, +1, +1, +1,
				-2, -4, -3, -2, -3, -4, -4, -1,
				-1, -1, -4, -1, -1, -2, -2, -3,
				+4, +3, +1, +4, +1, +4, +3, +3,
				-4, -3, -3, -3, -4, -3, -4, -1,
				+4, +3, +4, +1, +1, +3, +1, +3,
				-2, -4, -3, -1, -1, -3, -4, -2,
				+4, +4, +2, +4, +4, +1, +3, +3,
				-2, -1, -4, -1, -2, -4, -3, -4,
				+1, +1, +3, +1, +1, +2, +1, +3,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var bef []float32
			{
				bef = tc.mod.States()
			}

			if !reflect.DeepEqual(bef, tc.bef) {
				t.Fatalf("bef\n\n%s\n", cmp.Diff(tc.bef, bef))
			}

			{
				tc.mod.Update(tc.inp)
			}

			var aft []float32
			{
				aft = tc.mod.States()
			}

			if !reflect.DeepEqual(aft, tc.aft) {
				t.Fatalf("aft\n\n%s\n", cmp.Diff(tc.aft, aft))
			}
		})
	}
}
