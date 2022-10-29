package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/module/static"
	"github.com/phoebetron/getlin/vector"
)

func Test_Mapper_Module_Search(t *testing.T) {
	var mo1 getlin.Module
	var mo2 getlin.Module
	var mo3 getlin.Module
	var mo4 getlin.Module
	var mo5 getlin.Module
	var mo6 getlin.Module
	var mo7 getlin.Module
	{
		mo1 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo2 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo3 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo4 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo5 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo6 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo7 = static.New(static.Config{Out: vector.New(vector.Config{})})
	}

	testCases := []struct {
		mod getlin.Module
		ind int
		abo []getlin.Module
		bel []getlin.Module
	}{
		// Case 0
		{
			mod: mo1,
			ind: 0,
			abo: nil,
			bel: []getlin.Module{
				mo3,
				mo4,
				mo5,
				mo6,
			},
		},
		// Case 1
		{
			mod: mo2,
			ind: 1,
			abo: nil,
			bel: []getlin.Module{
				mo3,
				mo4,
				mo5,
				mo6,
			},
		},
		// Case 2
		{
			mod: mo3,
			ind: 2,
			abo: []getlin.Module{
				mo1,
				mo2,
			},
			bel: []getlin.Module{
				mo7,
			},
		},
		// Case 3
		{
			mod: mo5,
			ind: 4,
			abo: []getlin.Module{
				mo1,
				mo2,
			},
			bel: []getlin.Module{
				mo7,
			},
		},
		// Case 4
		{
			mod: mo7,
			ind: 6,
			abo: []getlin.Module{
				mo3,
				mo4,
				mo5,
				mo6,
			},
			bel: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var mpr *Mapper
			{
				mpr = New(Config{
					Mod: [][]getlin.Module{
						{
							mo1,
							mo2,
						},
						{
							mo3,
							mo4,
							mo5,
							mo6,
						},
						{
							mo7,
						},
					},
				})
			}

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

			if !reflect.DeepEqual(abo, tc.abo) {
				t.Fatalf("modules above do not match")
			}
			if !reflect.DeepEqual(bel, tc.bel) {
				t.Fatalf("modules below do not match")
			}
			if !reflect.DeepEqual(ind, tc.ind) {
				t.Fatalf("ind\n\n%s\n", cmp.Diff(tc.ind, ind))
			}
		})
	}
}
