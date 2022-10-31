package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Vector_Binary_Spl(t *testing.T) {
	testCases := []struct {
		raw []int
		ind int
		lef []int
		rig []int
	}{
		// case 0
		{
			raw: nil,
			ind: 0,
			lef: []int{},
			rig: []int{},
		},
		// case 1
		{
			raw: []int{
				0,
				0,
				1,
				0,
				1,
			},
			ind: 0,
			lef: []int{},
			rig: []int{
				0,
				0,
				1,
				0,
				1,
			},
		},
		// case 2
		{
			raw: []int{
				0,
				0,
				1,
				0,
				1,
			},
			ind: 1,
			lef: []int{
				0,
			},
			rig: []int{
				0,
				1,
				0,
				1,
			},
		},
		// case 3
		{
			raw: []int{
				0,
				0,
				1,
				0,
				1,
			},
			ind: 2,
			lef: []int{
				0,
				0,
			},
			rig: []int{
				1,
				0,
				1,
			},
		},
		// case 4
		{
			raw: []int{
				0,
				0,
				1,
				0,
				1,
			},
			ind: 3,
			lef: []int{
				0,
				0,
				1,
			},
			rig: []int{
				0,
				1,
			},
		},
		// case 5
		{
			raw: []int{
				0,
				0,
				1,
				0,
				1,
			},
			ind: 4,
			lef: []int{
				0,
				0,
				1,
				0,
			},
			rig: []int{
				1,
			},
		},
		// case 6
		{
			raw: []int{
				0,
				0,
				1,
				0,
				1,
			},
			ind: 5,
			lef: []int{
				0,
				0,
				1,
				0,
				1,
			},
			rig: nil,
		},
		// case 7
		{
			raw: []int{
				0,
				0,
				1,
				0,
				1,
			},
			ind: 6,
			lef: []int{
				0,
				0,
				1,
				0,
				1,
			},
			rig: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var raw getlin.Binary
			{
				raw = newbin(tobool(tc.raw)...)
			}

			var lef getlin.Binary
			var rig getlin.Binary
			{
				lef, rig = raw.Spl(tc.ind)
			}

			if !reflect.DeepEqual(tobool(tc.lef), lef.Raw()) {
				t.Fatalf("lef\n\n%s\n", cmp.Diff(tobool(tc.lef), lef.Raw()))
			}
			if !reflect.DeepEqual(tobool(tc.rig), rig.Raw()) {
				t.Fatalf("rig\n\n%s\n", cmp.Diff(tobool(tc.rig), rig.Raw()))
			}
		})
	}
}

func tobool(lis []int) []bool {
	bol := []bool{}

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
