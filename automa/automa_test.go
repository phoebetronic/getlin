package automa

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Automa_Lifecycle(t *testing.T) {
	testCases := []struct {
		set func(Interface)
		sta int
		poi float32
		rat float32
	}{
		// Case 0
		{
			set: func(a Interface) {
				// nothing
			},
			sta: 4,
			poi: 0,
			rat: 0,
		},
		// Case 1
		{
			set: func(a Interface) {
				a.Add(1)
				a.Add(1)
			},
			sta: 4,
			poi: 2,
			rat: 0.5,
		},
		// Case 2
		{
			set: func(a Interface) {
				a.Add(3)
				a.Add(1)
			},
			sta: 4,
			poi: 4,
			rat: 1.0,
		},
		// Case 3
		{
			set: func(a Interface) {
				a.Add(3)
				a.Add(2)
				a.Add(5)
			},
			sta: 4,
			poi: 4,
			rat: 1.0,
		},
		// Case 4
		{
			set: func(a Interface) {
				a.Add(3)
				a.Add(5)
				a.Rem(2)
			},
			sta: 4,
			poi: 2,
			rat: 0.5,
		},
		// Case 5
		{
			set: func(a Interface) {
				a.Add(3)
				a.Add(5)
				a.Rem(16)
			},
			sta: 4,
			poi: -4,
			rat: 1.0,
		},
		// Case 6
		{
			set: func(a Interface) {
				a.Rem(3)
				a.Add(5)
				a.Rem(3)
			},
			sta: 4,
			poi: -1,
			rat: 0.25,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var a Interface
			{
				a = New(Config{
					Sta: tc.sta,
				})
			}

			{
				tc.set(a)
			}

			var poi float32
			{
				poi = a.Poi()
			}

			var rat float32
			{
				rat = a.Rat()
			}

			if !reflect.DeepEqual(poi, tc.poi) {
				t.Fatalf("poi\n\n%s\n", cmp.Diff(tc.poi, poi))
			}
			if !reflect.DeepEqual(rat, tc.rat) {
				t.Fatalf("rat\n\n%s\n", cmp.Diff(tc.rat, rat))
			}
		})
	}
}
