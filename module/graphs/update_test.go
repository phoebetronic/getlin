package graphs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/module/linear"
	"github.com/phoebetron/getlin/random/static"
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
			mod: New(Config{
				Mod: [][]getlin.Module{
					{
						muslin(2, 1),
						muslin(2, 1),
					},
					{
						muslin(2, 2),
					},
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
				0.125,
				0,
				0,
				0,
				0,
				0.125,
				0,
				0,
				0,
				0, // -
				0, // 0
				0, // +
				0,
				0,
				0,
				0.125,
				0,
				0,
				0,
				0,
				0.125,
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
			mod: New(Config{
				Mod: [][]getlin.Module{
					{
						muslin(2, 1),
						muslin(2, 1),
					},
					{
						muslin(2, 2),
					},
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
				-3,
				-2,
				+4,
				+1,
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

func muslin(inp int, out int) getlin.Module {
	return linear.New(linear.Config{
		Inp: inp,
		Out: out,
		Ran: static.New(static.Config{F32: 0.05}),
		Sta: 4,
	})
}
