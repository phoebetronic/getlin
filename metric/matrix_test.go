package metric

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Metric_Matrix(t *testing.T) {
	testCases := []struct {
		sam [4]int
		acc float32
		mae float32
	}{
		// Case 0
		{
			sam: [4]int{
				TP: 3,
				TN: 0,
				FP: 0,
				FN: 0,
			},
			acc: 1,
			mae: 0,
		},
		// Case 1
		{
			sam: [4]int{
				TP: 0,
				TN: 3,
				FP: 0,
				FN: 0,
			},
			acc: 1,
			mae: 0,
		},
		// Case 2
		{
			sam: [4]int{
				TP: 6,
				TN: 3,
				FP: 0,
				FN: 0,
			},
			acc: 1,
			mae: 0,
		},
		// Case 3
		{
			sam: [4]int{
				TP: 0,
				TN: 0,
				FP: 5,
				FN: 0,
			},
			acc: 0,
			mae: 1,
		},
		// Case 4
		{
			sam: [4]int{
				TP: 0,
				TN: 0,
				FP: 0,
				FN: 5,
			},
			acc: 0,
			mae: 1,
		},
		// Case 5
		{
			sam: [4]int{
				TP: 0,
				TN: 0,
				FP: 2,
				FN: 4,
			},
			acc: 0,
			mae: 1,
		},
		// Case 6
		{
			sam: [4]int{
				TP: 4,
				TN: 0,
				FP: 2,
				FN: 0,
			},
			acc: 0.6666666865348816,
			mae: 0.3333333432674408,
		},
		// Case 7
		{
			sam: [4]int{
				TP: 5,
				TN: 8,
				FP: 2,
				FN: 1,
			},
			acc: 0.8125,
			mae: 0.1875,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var met getlin.Metric
			{
				met = New(Config{})
			}

			for i, x := range tc.sam {
				for j := 0; j < x; j++ {
					met.Set().Mat(i, 1)
				}
			}

			var acc float32
			{
				acc = met.Get().Mat().Acc()
			}

			var mae float32
			{
				mae = met.Get().Mat().Mae()
			}

			if !reflect.DeepEqual(acc, tc.acc) {
				t.Fatalf("acc\n\n%s\n", cmp.Diff(tc.acc, acc))
			}
			if !reflect.DeepEqual(mae, tc.mae) {
				t.Fatalf("mae\n\n%s\n", cmp.Diff(tc.mae, mae))
			}
		})
	}
}
