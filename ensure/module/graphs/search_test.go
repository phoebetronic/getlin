package graphs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Module_Graphs_Search(t *testing.T) {
	testCases := []struct {
		vec getlin.Vector
		mod getlin.Module
		res getlin.Vector
	}{
		// Case 0
		{
			vec: musvec([]int{0, 1, 0}, []int{1}),
			mod: musgra([][]getlin.Module{
				{
					muslin(3, 1),
					muslin(3, 1),
					muslin(3, 1),
				},
				{
					muslin(3, 1),
				},
			}),
			res: musvec([]int{0, 0, 0}, []int{0}),
		},
		// Case 1
		{
			vec: musvec([]int{0, 0, 1}, []int{0}),
			mod: musgra([][]getlin.Module{
				{
					muslin(3, 1),
					muslin(3, 1),
				},
				{
					muslin(2, 1),
					muslin(2, 1),
				},
				{
					muslin(2, 1),
				},
			}),
			res: musvec([]int{0, 0}, []int{0}),
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
