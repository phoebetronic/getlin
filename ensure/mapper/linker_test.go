package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
)

func Test_Mapper_Linker(t *testing.T) {
	var mo0 getlin.Module
	var mo1 getlin.Module
	var mo2 getlin.Module
	var mo3 getlin.Module
	var mo4 getlin.Module
	var mo5 getlin.Module
	var mo6 getlin.Module
	{
		mo0 = muslin(2, 4)
		mo1 = muslin(2, 4)

		mo2 = muslin(8, 2)
		mo3 = muslin(8, 2)
		mo4 = muslin(8, 2)
		mo5 = muslin(8, 2)

		mo6 = muslin(8, 1)
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
				},
			},
		})
	}

	testCases := []struct {
		mod getlin.Module
		ind int
		lay int
		abo []getlin.Module
		bel []getlin.Module
	}{
		// Case 0
		{
			mod: mo0,
			ind: 0,
			lay: 0,
			abo: nil,
			bel: []getlin.Module{
				mo2,
				mo3,
				mo4,
				mo5,
			},
		},
		// Case 1
		{
			mod: mo1,
			ind: 1,
			lay: 0,
			abo: nil,
			bel: []getlin.Module{
				mo2,
				mo3,
				mo4,
				mo5,
			},
		},
		// Case 2
		{
			mod: mo2,
			ind: 2,
			lay: 1,
			abo: []getlin.Module{
				mo0,
				mo1,
			},
			bel: []getlin.Module{
				mo6,
			},
		},
		// Case 3
		{
			mod: mo4,
			ind: 4,
			lay: 1,
			abo: []getlin.Module{
				mo0,
				mo1,
			},
			bel: []getlin.Module{
				mo6,
			},
		},
		// Case 4
		{
			mod: mo6,
			ind: 6,
			lay: 2,
			abo: []getlin.Module{
				mo2,
				mo3,
				mo4,
				mo5,
			},
			bel: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var abo []getlin.Module
			{
				abo = mpr.Abo(tc.mod)
			}

			var bel []getlin.Module
			{
				bel = mpr.Bel(tc.mod)
			}

			var ind int
			{
				ind = mpr.Ind(tc.mod)
			}

			var lay int
			{
				lay = mpr.Lay(tc.mod)
			}

			if !reflect.DeepEqual(abo, tc.abo) {
				t.Fatalf("modules above do not match")
			}
			if !reflect.DeepEqual(bel, tc.bel) {
				t.Fatalf("modules below do not match")
			}
			if !reflect.DeepEqual(ind, tc.ind) {
				t.Fatalf("ind\n\n%s\n", cmp.Diff(tc.ind, ind))
			}
			if !reflect.DeepEqual(lay, tc.lay) {
				t.Fatalf("lay\n\n%s\n", cmp.Diff(tc.lay, lay))
			}
		})
	}
}
