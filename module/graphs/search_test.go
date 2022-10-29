package graphs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/module/invert"
	"github.com/phoebetron/getlin/module/static"
	"github.com/phoebetron/getlin/vector"
)

func Test_Module_Graphs_Search(t *testing.T) {
	testCases := []struct {
		inp getlin.Vector
		mod getlin.Module
		out getlin.Vector
	}{
		// Case 0
		{
			inp: musvec([]int{}, []int{1}),
			mod: New(Config{
				Mod: [][]getlin.Module{
					{
						static.New(static.Config{Out: musvec([]int{0, 0, 1}, []int{0})}),
						static.New(static.Config{Out: musvec([]int{0, 0, 1}, []int{1})}),
						static.New(static.Config{Out: musvec([]int{0, 0, 1}, []int{0})}),
					},
				},
			}),
			out: musvec([]int{0, 0, 1, 0, 0, 1, 0, 0, 1}, []int{}),
		},
		// Case 1
		{
			inp: musvec([]int{0, 0, 1}, []int{1}),
			mod: New(Config{
				Mod: [][]getlin.Module{
					// The first layer receives the input [0 0 1] and creates
					// the output [1 0 0 1 0 0] using its two invert modules.
					{
						invert.New(invert.Config{}),
						invert.New(invert.Config{}),
					},
					// The second layer receives the input [1 0 0 1 0 0] and
					// creates the output [0 0 1 0 0 1 1 1 1] using its invert
					// and static module.
					{
						invert.New(invert.Config{}),
						static.New(static.Config{Out: musvec([]int{1, 1, 1}, []int{1})}),
					},
					// The third layer receives the input [0 0 1 0 0 1 1 1 1]
					// and creates the output [1 1 1 1 0 0 1 0 0] using its
					// invert module.
					{
						invert.New(invert.Config{}),
					},
				},
			}),
			out: musvec([]int{1, 1, 1, 1, 0, 0, 1, 0, 0}, []int{}),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var out getlin.Vector
			{
				out = tc.mod.Search(tc.inp)
			}

			if !reflect.DeepEqual(out.Bit(), tc.out.Bit()) {
				t.Fatalf("bit\n\n%s\n", cmp.Diff(tc.out.Bit(), out.Bit()))
			}
			if !reflect.DeepEqual(out.Tru(), tc.out.Tru()) {
				t.Fatalf("tru\n\n%s\n", cmp.Diff(tc.out.Tru(), out.Tru()))
			}
		})
	}
}

func musvec(bit []int, tru []int) getlin.Vector {
	return vector.New(vector.Config{
		Bit: tobool(bit),
		Tru: tobool(tru),
	})
}

func tobool(lis []int) []bool {
	var bol []bool

	for _, x := range lis {
		if x == 0 {
			bol = append(bol, false)
		}
		if x == 1 {
			bol = append(bol, true)
		}
	}

	return bol
}
