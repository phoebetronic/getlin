package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
)

func Test_Mapper_Tru_003(t *testing.T) {
	var mo0 getlin.Module
	var mo1 getlin.Module
	var mo2 getlin.Module
	var mo3 getlin.Module
	var mo4 getlin.Module
	var mo5 getlin.Module
	var mo6 getlin.Module
	var mo7 getlin.Module
	{
		mo0 = muslin(3, 2)
		mo1 = muslin(7, 6)

		mo2 = muslin(8, 2)
		mo3 = muslin(8, 2)
		mo4 = muslin(8, 2)
		mo5 = muslin(8, 2)

		mo6 = muslin(8, 6)
		mo7 = muslin(8, 2)
	}

	var mpr getlin.Mapper
	{
		mpr = mapper.New(mapper.Config{
			Mod: [][]getlin.Module{
				{
					mo0,
					mo1,
				},
				{
					mo2,
					mo3,
					mo4,
					mo5,
				},
				{
					mo6,
					mo7,
				},
			},
		})
	}

	testCases := []struct {
		mod getlin.Module
		tru [2]int
	}{
		// Case 0
		{
			mod: mo0,
			tru: [2]int{0, 2},
		},
		// Case 1
		{
			mod: mo1,
			tru: [2]int{2, 8},
		},
		// Case 2
		{
			mod: mo2,
			tru: [2]int{0, 2},
		},
		// Case 3
		{
			mod: mo3,
			tru: [2]int{2, 4},
		},
		// Case 4
		{
			mod: mo4,
			tru: [2]int{4, 6},
		},
		// Case 5
		{
			mod: mo5,
			tru: [2]int{6, 8},
		},
		// Case 6
		{
			mod: mo6,
			tru: [2]int{0, 6},
		},
		// Case 7
		{
			mod: mo7,
			tru: [2]int{6, 8},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var tru [2]int
			{
				tru = mpr.Tru(tc.mod)
			}

			if !reflect.DeepEqual(tru, tc.tru) {
				t.Fatalf("tru\n\n%s\n", cmp.Diff(tc.tru, tru))
			}
		})
	}
}
