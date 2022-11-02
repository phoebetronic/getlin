package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Vector_Binary_Maj(t *testing.T) {
	testCases := []struct {
		raw []int
		maj int
	}{
		// Case 0
		{
			raw: nil,
			maj: 0,
		},
		// Case 1
		{
			raw: []int{
				0,
			},
			maj: 0,
		},
		// Case 2
		{
			raw: []int{
				0,
				0,
				0,
			},
			maj: 0,
		},
		// Case 3
		{
			raw: []int{
				0,
				0,
				1,
				0,
			},
			maj: 0,
		},
		// Case 4
		{
			raw: []int{
				0,
				0,
				1,
				1,
				0,
				1,
			},
			maj: 0,
		},
		// Case 5
		{
			raw: []int{
				0,
				1,
				0,
				1,
				1,
				0,
				1,
			},
			maj: 1,
		},
		// Case 6
		{
			raw: []int{
				0,
				1,
				1,
				1,
				1,
			},
			maj: 1,
		},
		// Case 7
		{
			raw: []int{
				1,
			},
			maj: 1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var raw getlin.Binary
			{
				raw = newbin(tobool(tc.raw...)...)
			}

			var maj []int
			{
				maj = tobits(raw.Maj())
			}

			if !reflect.DeepEqual([]int{tc.maj}, maj) {
				t.Fatalf("maj\n\n%s\n", cmp.Diff([]int{tc.maj}, maj))
			}
		})
	}
}

func Test_Vector_Binary_Spl(t *testing.T) {
	testCases := []struct {
		raw []int
		ind int
		lef []int
		rig []int
	}{
		// Case 0
		{
			raw: nil,
			ind: 0,
			lef: []int{},
			rig: []int{},
		},
		// Case 1
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
		// Case 2
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
		// Case 3
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
		// Case 4
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
		// Case 5
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
		// Case 6
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
		// Case 7
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
				raw = newbin(tobool(tc.raw...)...)
			}

			var lef getlin.Binary
			var rig getlin.Binary
			{
				lef, rig = raw.Spl(tc.ind)
			}

			if !reflect.DeepEqual(tobool(tc.lef...), lef.Raw()) {
				t.Fatalf("lef\n\n%s\n", cmp.Diff(tobool(tc.lef...), lef.Raw()))
			}
			if !reflect.DeepEqual(tobool(tc.rig...), rig.Raw()) {
				t.Fatalf("rig\n\n%s\n", cmp.Diff(tobool(tc.rig...), rig.Raw()))
			}
		})
	}
}

func Test_Vector_Binary_Wei(t *testing.T) {
	testCases := []struct {
		raw []int
		bit int
		wei float32
	}{
		// Case 0
		{
			raw: []int{
				0,
			},
			bit: 0,
			wei: 1.0,
		},
		// Case 1
		{
			raw: []int{
				0,
			},
			bit: 1,
			wei: 0.0,
		},
		// Case 2
		{
			raw: []int{
				0,
				0,
				0,
			},
			bit: 0,
			wei: 1.0,
		},
		// Case 3
		{
			raw: []int{
				0,
				0,
				0,
			},
			bit: 1,
			wei: 0.0,
		},
		// Case 4
		{
			raw: []int{
				1,
			},
			bit: 0,
			wei: 0.0,
		},
		// Case 5
		{
			raw: []int{
				1,
			},
			bit: 1,
			wei: 1.0,
		},
		// Case 6
		{
			raw: []int{
				1,
				1,
				1,
			},
			bit: 0,
			wei: 0.0,
		},
		// Case 7
		{
			raw: []int{
				1,
				1,
				1,
			},
			bit: 1,
			wei: 1.0,
		},
		// Case 8
		{
			raw: []int{
				0,
				1,
				0,
				0,
				1,
				1,
				1,
				0,
				1,
			},
			bit: 0,
			wei: 0.4444444477558136,
		},
		// Case 9
		{
			raw: []int{
				0,
				1,
				0,
				0,
				1,
				1,
				1,
				0,
				1,
			},
			bit: 1,
			wei: 0.5555555820465088,
		},
		// Case 10
		{
			raw: []int{
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				0,
				1,
				1,
				1,
				1,
				1,
			},
			bit: 0,
			wei: 0.2222222238779068,
		},
		// Case 11
		{
			raw: []int{
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				0,
				1,
				1,
				1,
				1,
				1,
			},
			bit: 1,
			wei: 0.7777777910232544,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var raw getlin.Binary
			{
				raw = newbin(tobool(tc.raw...)...)
			}

			var wei float32
			{
				wei = raw.Wei(tobool(tc.bit)[0])
			}

			if !reflect.DeepEqual(tc.wei, wei) {
				t.Fatalf("wei\n\n%s\n", cmp.Diff(tc.wei, wei))
			}
		})
	}
}

func tobits(raw ...bool) []int {
	var bit []int

	for _, x := range raw {
		if x {
			bit = append(bit, 1)
		} else {
			bit = append(bit, 0)
		}
	}

	return bit
}

func tobool(lis ...int) []bool {
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
