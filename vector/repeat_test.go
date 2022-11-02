package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Vector_Repeat(t *testing.T) {
	testCases := []struct {
		bit []bool
		cou int
		rep []bool
	}{
		// Case 0
		{
			bit: nil,
			cou: 0,
			rep: nil,
		},
		// Case 1
		{
			bit: []bool{
				true,
			},
			cou: 1,
			rep: []bool{
				true,
			},
		},
		// Case 2
		{
			bit: []bool{
				false,
			},
			cou: 1,
			rep: []bool{
				false,
			},
		},
		// Case 3
		{
			bit: []bool{
				true,
			},
			cou: 3,
			rep: []bool{
				true,
				true,
				true,
			},
		},
		// Case 4
		{
			bit: []bool{
				false,
			},
			cou: 3,
			rep: []bool{
				false,
				false,
				false,
			},
		},
		// Case 5
		{
			bit: []bool{
				true,
				false,
			},
			cou: 1,
			rep: []bool{
				true,
				false,
			},
		},
		// Case 6
		{
			bit: []bool{
				true,
				false,
			},
			cou: 4,
			rep: []bool{
				true,
				false,
				true,
				false,
				true,
				false,
				true,
				false,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var rep []bool
			{
				rep = Repeat(tc.bit, tc.cou)
			}

			if !reflect.DeepEqual(tc.rep, rep) {
				t.Fatalf("rep\n\n%s\n", cmp.Diff(tc.rep, rep))
			}
		})
	}
}
