package automa

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Automa_Lifecycle(t *testing.T) {
	testCases := []struct {
		set func(getlin.Automa)
		sta int
		poi int
		inc bool
		exc bool
		hig bool
		low bool
	}{
		// Case 0
		{
			set: func(a getlin.Automa) {
				// nothing
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: false,
		},
		// Case 1
		{
			set: func(a getlin.Automa) {
				a.Add()
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: false,
			low: false,
		},
		// Case 2
		{
			set: func(a getlin.Automa) {
				a.Add()
				a.Add()
				a.Red()
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: false,
			low: false,
		},
		// Case 3
		{
			set: func(a getlin.Automa) {
				a.Red()
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: false,
		},
		// Case 4
		{
			set: func(a getlin.Automa) {
				a.Red()
				a.Red()
				a.Add()
				a.Add()
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: false,
		},
		// Case 5
		{
			set: func(a getlin.Automa) {
				a.Red()
				a.Red()
				a.Add()
				a.Add()
				a.Add()
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: false,
			low: false,
		},
		// Case 6
		{
			set: func(a getlin.Automa) {
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: true,
		},
		// Case 7
		{
			set: func(a getlin.Automa) {
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()
				a.Red()

				a.Add()
				a.Add()
				a.Add()
				a.Add()
				a.Add()
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: false,
			low: false,
		},
		// Case 8
		{
			set: func(a getlin.Automa) {
				a.Add() // 5
				a.Add() // 6
				a.Add() // 7
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: false,
			low: false,
		},
		// Case 9
		{
			set: func(a getlin.Automa) {
				a.Add() // 5
				a.Add() // 6
				a.Add() // 7
				a.Add() // 8
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: true,
			low: false,
		},
		// Case 10
		{
			set: func(a getlin.Automa) {
				a.Add() // 5
				a.Add() // 6
				a.Add() // 7
				a.Add() // 8
				a.Add() // 8
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: true,
			low: false,
		},
		// Case 11
		{
			set: func(a getlin.Automa) {
				a.Red() // 3
				a.Red() // 2
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: false,
		},
		// Case 12
		{
			set: func(a getlin.Automa) {
				a.Red() // 3
				a.Red() // 2
				a.Red() // 1
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: true,
		},
		// Case 13
		{
			set: func(a getlin.Automa) {
				a.Red() // 3
				a.Red() // 2
				a.Red() // 1
				a.Red() // 1
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: true,
		},
		// Case 14
		{
			set: func(a getlin.Automa) {
				a.Add() // 5
				a.Add() // 6
				a.Add() // 7
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8

				a.Red() // 7
				a.Red() // 6
				a.Red() // 5
			},
			sta: 4,
			poi: 4,
			inc: true,
			exc: false,
			hig: false,
			low: false,
		},
		// Case 15
		{
			set: func(a getlin.Automa) {
				a.Add() // 5
				a.Add() // 6
				a.Add() // 7
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8
				a.Add() // 8

				a.Red() // 7
				a.Red() // 6
				a.Red() // 5
				a.Red() // 4
			},
			sta: 4,
			poi: 4,
			inc: false,
			exc: true,
			hig: false,
			low: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var a getlin.Automa
			{
				a = New(Config{
					Poi: tc.poi,
					Sta: tc.sta,
				})
			}

			{
				tc.set(a)
			}

			var inc bool
			{
				inc = a.Inc()
			}

			var exc bool
			{
				exc = a.Exc()
			}

			var hig bool
			{
				hig = a.Hig()
			}

			var low bool
			{
				low = a.Low()
			}

			if !reflect.DeepEqual(inc, tc.inc) {
				t.Fatalf("inc\n\n%s\n", cmp.Diff(tc.inc, inc))
			}
			if !reflect.DeepEqual(exc, tc.exc) {
				t.Fatalf("exc\n\n%s\n", cmp.Diff(tc.exc, exc))
			}
			if !reflect.DeepEqual(hig, tc.hig) {
				t.Fatalf("hig\n\n%s\n", cmp.Diff(tc.hig, hig))
			}
			if !reflect.DeepEqual(low, tc.low) {
				t.Fatalf("low\n\n%s\n", cmp.Diff(tc.low, low))
			}
		})
	}
}
