package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
)

func Test_Mapper_Tru_005(t *testing.T) {
	var mo0 getlin.Module
	var mo1 getlin.Module
	var mo2 getlin.Module
	var mo3 getlin.Module
	var mo4 getlin.Module
	var mo5 getlin.Module
	{
		mo0 = muslin(32, 16)
		mo1 = muslin(32, 16)
		mo2 = muslin(32, 16)
		mo3 = muslin(32, 16)

		mo4 = musvot(64)
		mo5 = musvot(64)
	}

	var mpr getlin.Mapper
	{
		mpr = mapper.New(mapper.Config{
			Mod: [][]getlin.Module{
				{
					mo0,
					mo1,
					mo2,
					mo3,
				},
				{
					mo4,
					mo5,
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
			tru: [2]int{0, 1},
		},
		// Case 1
		{
			mod: mo1,
			tru: [2]int{0, 1},
		},
		// Case 2
		{
			mod: mo2,
			tru: [2]int{1, 2},
		},
		// Case 3
		{
			mod: mo3,
			tru: [2]int{1, 2},
		},
		// Case 4
		{
			mod: mo4,
			tru: [2]int{0, 1},
		},
		// Case 5
		{
			mod: mo5,
			tru: [2]int{1, 2},
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
