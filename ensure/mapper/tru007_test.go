package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
)

func Test_Mapper_Tru_007(t *testing.T) {
	var mo0 getlin.Module
	var mo1 getlin.Module
	var mo2 getlin.Module
	var mo3 getlin.Module
	var mo4 getlin.Module
	var mo5 getlin.Module
	var mo6 getlin.Module
	var mo7 getlin.Module
	var mo8 getlin.Module
	var mo9 getlin.Module
	{
		mo0 = muslin(3, 3)
		mo1 = muslin(3, 3)

		mo3 = muslin(6, 4)
		mo4 = muslin(6, 4)
		mo5 = musvot(8)
		mo2 = musgra([][]getlin.Module{
			{
				mo3,
				mo4,
			},
			{
				mo5,
			},
		})

		mo7 = muslin(6, 4)
		mo8 = muslin(6, 4)
		mo9 = musvot(8)
		mo6 = musgra([][]getlin.Module{
			{
				mo7,
				mo8,
			},
			{
				mo9,
			},
		})
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
					mo6,
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
			tru: [2]int{1, 2},
		},
		// Case 2
		{
			mod: mo2,
			tru: [2]int{0, 1},
		},
		// Case 6
		{
			mod: mo6,
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
