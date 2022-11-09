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
		ini float32
		inc bool
		exc bool
		hig bool
		low bool
		poi float32
	}{
		// Case 0
		{
			set: func(a getlin.Automa) {
				// nothing
			},
			ini: 0.5,
			inc: false,
			exc: true,
			hig: false,
			low: false,
			poi: 0.5,
		},
		// Case 1
		{
			set: func(a getlin.Automa) {
				a.Add(0.1) // 0.6
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: false,
			low: false,
			poi: 0.6000000238418579,
		},
		// Case 2
		{
			set: func(a getlin.Automa) {
				a.Add(0.1) // 0.6
				a.Add(0.1) // 0.7
				a.Red(0.1) // 0.6
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: false,
			low: false,
			poi: 0.6000000238418579,
		},
		// Case 3
		{
			set: func(a getlin.Automa) {
				a.Red(0.1) // 0.4
			},
			ini: 0.5,
			inc: false,
			exc: true,
			hig: false,
			low: false,
			poi: 0.4000000059604645,
		},
		// Case 4
		{
			set: func(a getlin.Automa) {
				a.Red(0.1) // 0.4
				a.Red(0.1) // 0.3
				a.Add(0.1) // 0.4
				a.Add(0.1) // 0.5
			},
			ini: 0.5,
			inc: false,
			exc: true,
			hig: false,
			low: false,
			poi: 0.5,
		},
		// Case 5
		{
			set: func(a getlin.Automa) {
				a.Red(0.1) // 0.4
				a.Red(0.1) // 0.3
				a.Add(0.1) // 0.4
				a.Add(0.1) // 0.5
				a.Add(0.1) // 0.6
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: false,
			low: false,
			poi: 0.6000000238418579,
		},
		// Case 6
		{
			set: func(a getlin.Automa) {
				a.Red(0.1) // 0.4
				a.Red(0.1) // 0.3
				a.Red(0.1) // 0.2
				a.Red(0.1) // 0.1
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
			},
			ini: 0.5,
			inc: false,
			exc: true,
			hig: false,
			low: true,
			poi: 0.0,
		},
		// Case 7
		{
			set: func(a getlin.Automa) {
				a.Red(0.1) // 0.4
				a.Red(0.1) // 0.3
				a.Red(0.1) // 0.2
				a.Red(0.1) // 0.1
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0

				a.Add(0.1) // 0.1
				a.Add(0.1) // 0.2
				a.Add(0.1) // 0.3
				a.Add(0.1) // 0.4
				a.Add(0.1) // 0.5
				a.Add(0.1) // 0.6
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: false,
			low: false,
			poi: 0.6000000238418579,
		},
		// Case 8
		{
			set: func(a getlin.Automa) {
				a.Add(0.1) // 0.6
				a.Add(0.1) // 0.7
				a.Add(0.1) // 0.8
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: false,
			low: false,
			poi: 0.8000000715255737,
		},
		// Case 9
		{
			set: func(a getlin.Automa) {
				a.Add(0.1) // 0.6
				a.Add(0.1) // 0.7
				a.Add(0.1) // 0.8
				a.Add(0.1) // 0.9
				a.Add(0.1) // 1.0
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: true,
			low: false,
			poi: 1.0,
		},
		// Case 10
		{
			set: func(a getlin.Automa) {
				a.Add(0.1) // 0.6
				a.Add(0.1) // 0.7
				a.Add(0.1) // 0.8
				a.Add(0.1) // 0.9
				a.Add(0.1) // 1.0
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: true,
			low: false,
			poi: 1.0,
		},
		// Case 11
		{
			set: func(a getlin.Automa) {
				a.Red(0.1) // 0.4
				a.Red(0.1) // 0.3
			},
			ini: 0.5,
			inc: false,
			exc: true,
			hig: false,
			low: false,
			poi: 0.30000001192092896,
		},
		// Case 12
		{
			set: func(a getlin.Automa) {
				a.Red(0.1) // 0.4
				a.Red(0.1) // 0.3
				a.Red(0.1) // 0.2
				a.Red(0.1) // 0.1
				a.Red(0.1) // 0.0
				a.Red(0.1) // 0.0
			},
			ini: 0.5,
			inc: false,
			exc: true,
			hig: false,
			low: true,
			poi: 0.0,
		},
		// Case 13
		{
			set: func(a getlin.Automa) {
				a.Add(0.1) // 0.6
				a.Add(0.1) // 0.7
				a.Add(0.1) // 0.8
				a.Add(0.1) // 0.9
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0

				a.Red(0.1) // 0.9
				a.Red(0.1) // 0.8
				a.Red(0.1) // 0.7
			},
			ini: 0.5,
			inc: true,
			exc: false,
			hig: false,
			low: false,
			poi: 0.6999999284744263,
		},
		// Case 14
		{
			set: func(a getlin.Automa) {
				a.Add(0.1) // 0.6
				a.Add(0.1) // 0.7
				a.Add(0.1) // 0.8
				a.Add(0.1) // 0.9
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0
				a.Add(0.1) // 1.0

				a.Red(0.1) // 0.9
				a.Red(0.1) // 0.8
				a.Red(0.1) // 0.7
				a.Red(0.1) // 0.6
				a.Red(0.1) // 0.5
			},
			ini: 0.5,
			inc: false,
			exc: true,
			hig: false,
			low: false,
			poi: 0.49999991059303284,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var a getlin.Automa
			{
				a = New(Config{
					Poi: tc.ini,
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

			var poi float32
			{
				poi = a.Poi()
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
			if !reflect.DeepEqual(poi, tc.poi) {
				t.Fatalf("poi\n\n%s\n", cmp.Diff(tc.poi, poi))
			}
		})
	}
}
