package metric

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Metric_Errors(t *testing.T) {
	testCases := []struct {
		set [][2]float32
		mae float32
	}{
		// Case 0
		{
			set: [][2]float32{
				{1.0, 1.0},
				{0.2, 0.2},
			},
			mae: 0.0,
		},
		// Case 1
		{
			set: [][2]float32{
				{0.0, 0.0},
				{0.0, 0.0},
			},
			mae: 0.0,
		},
		// Case 2
		{
			set: [][2]float32{
				{0.0, 0.0},
				{1.0, 1.0},
				{0.0, 0.0},
				{0.5, 0.5},
			},
			mae: 0.0,
		},
		// Case 3
		{
			set: [][2]float32{
				{0.0, 1.0},
			},
			mae: 1.0,
		},
		// Case 4
		{
			set: [][2]float32{
				{0.5, 0.3},
			},
			mae: 0.19999998807907104,
		},
		// Case 5
		{
			set: [][2]float32{
				{0.0, 1.0},
				{0.5, 0.3},
			},
			mae: 0.6000000238418579,
		},
		// Case 6
		{
			set: [][2]float32{
				{0.0, 1.0},
				{1.0, 1.0},
				{0.5, 0.3},
				{0.3, 0.3},
				{0.6, 0.6},
				{0.3, 0.3},
				{0.7, 0.6},
				{0.9, 0.0},
				{0.9, 0.9},
				{0.4, 0.4},
				{0.4, 0.4},
			},
			mae: 0.19999998807907104,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var met getlin.Metric
			{
				met = New(Config{})
			}

			for _, x := range tc.set {
				met.Set().Err(x[0], x[1])
			}

			var mae float32
			{
				mae = met.Get().Err().Mae()
			}

			if !reflect.DeepEqual(mae, tc.mae) {
				t.Fatalf("mae\n\n%s\n", cmp.Diff(tc.mae, mae))
			}
		})
	}
}
