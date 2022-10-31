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
		vec getlin.Vector
		mod getlin.Module
		res getlin.Vector
	}{
		// Case 0
		{
			vec: musvec([]int{}, []int{1}),
			mod: New(Config{
				Mod: [][]getlin.Module{
					{
						static.New(static.Config{Out: musvec([]int{0}, []int{0, 0, 1})}),
						static.New(static.Config{Out: musvec([]int{1}, []int{0, 0, 1})}),
						static.New(static.Config{Out: musvec([]int{0}, []int{0, 0, 1})}),
					},
				},
			}),
			res: musvec([]int{}, []int{0, 0, 1, 0, 0, 1, 0, 0, 1}),
		},
		// Case 1
		{
			vec: musvec([]int{0, 0, 1}, []int{1}),
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
						static.New(static.Config{Out: musvec([]int{1}, []int{1, 1, 1})}),
					},
					// The third layer receives the input [0 0 1 0 0 1 1 1 1]
					// and creates the output [1 1 1 1 0 0 1 0 0] using its
					// invert module.
					{
						invert.New(invert.Config{}),
					},
				},
			}),
			res: musvec([]int{0, 0, 1, 0, 0, 1, 1, 1, 1}, []int{1, 1, 1, 1, 0, 0, 1, 0, 0}),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var vec getlin.Vector
			{
				vec = tc.mod.Search(tc.vec)
			}

			if !reflect.DeepEqual(vec.Inp().Raw(), tc.res.Inp().Raw()) {
				t.Fatalf("inp\n\n%s\n", cmp.Diff(tc.res.Inp().Raw(), vec.Inp().Raw()))
			}
			if !reflect.DeepEqual(vec.Out().Raw(), tc.res.Out().Raw()) {
				t.Fatalf("out\n\n%s\n", cmp.Diff(tc.res.Out().Raw(), vec.Out().Raw()))
			}
		})
	}
}

func musvec(inp []int, out []int) getlin.Vector {
	return vector.New(vector.Config{
		Inp: tobool(inp),
		Out: tobool(out),
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
