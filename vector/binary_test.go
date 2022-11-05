package vector

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
)

func Test_Vector_Binary_Maj(t *testing.T) {
	testCases := []struct {
		raw []uint8
		maj uint8
	}{
		// Case 0
		{
			raw: nil,
			maj: 0,
		},
		// Case 1
		{
			raw: []uint8{
				0,
			},
			maj: 0,
		},
		// Case 2
		{
			raw: []uint8{
				0,
				0,
				0,
			},
			maj: 0,
		},
		// Case 3
		{
			raw: []uint8{
				0,
				0,
				1,
				0,
			},
			maj: 0,
		},
		// Case 4
		{
			raw: []uint8{
				0,
				0,
				1,
				1,
				0,
				1,
			},
			maj: 0,
		},
		// Case 5
		{
			raw: []uint8{
				0,
				1,
				0,
				1,
				1,
				0,
				1,
			},
			maj: 1,
		},
		// Case 6
		{
			raw: []uint8{
				0,
				1,
				1,
				1,
				1,
			},
			maj: 1,
		},
		// Case 7
		{
			raw: []uint8{
				1,
			},
			maj: 1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var raw getlin.Binary
			{
				raw = newbin(tc.raw...)
			}

			var maj uint8
			{
				maj = raw.Maj()
			}

			if !reflect.DeepEqual(tc.maj, maj) {
				t.Fatalf("maj\n\n%s\n", cmp.Diff(tc.maj, maj))
			}
		})
	}
}

func Test_Vector_Binary_Wei(t *testing.T) {
	testCases := []struct {
		raw []uint8
		bit uint8
		wei float32
	}{
		// Case 0
		{
			raw: []uint8{
				0,
			},
			bit: 0,
			wei: 1.0,
		},
		// Case 1
		{
			raw: []uint8{
				0,
			},
			bit: 1,
			wei: 0.0,
		},
		// Case 2
		{
			raw: []uint8{
				0,
				0,
				0,
			},
			bit: 0,
			wei: 1.0,
		},
		// Case 3
		{
			raw: []uint8{
				0,
				0,
				0,
			},
			bit: 1,
			wei: 0.0,
		},
		// Case 4
		{
			raw: []uint8{
				1,
			},
			bit: 0,
			wei: 0.0,
		},
		// Case 5
		{
			raw: []uint8{
				1,
			},
			bit: 1,
			wei: 1.0,
		},
		// Case 6
		{
			raw: []uint8{
				1,
				1,
				1,
			},
			bit: 0,
			wei: 0.0,
		},
		// Case 7
		{
			raw: []uint8{
				1,
				1,
				1,
			},
			bit: 1,
			wei: 1.0,
		},
		// Case 8
		{
			raw: []uint8{
				0,
				1,
				0,
				0,
				1,
				1,
				1,
				0,
				1,
			},
			bit: 0,
			wei: 0.4444444477558136,
		},
		// Case 9
		{
			raw: []uint8{
				0,
				1,
				0,
				0,
				1,
				1,
				1,
				0,
				1,
			},
			bit: 1,
			wei: 0.5555555820465088,
		},
		// Case 10
		{
			raw: []uint8{
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				0,
				1,
				1,
				1,
				1,
				1,
			},
			bit: 0,
			wei: 0.2222222238779068,
		},
		// Case 11
		{
			raw: []uint8{
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				1,
				1,
				1,
				0,
				0,
				1,
				1,
				1,
				1,
				1,
			},
			bit: 1,
			wei: 0.7777777910232544,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var raw getlin.Binary
			{
				raw = newbin(tc.raw...)
			}

			var wei float32
			{
				wei = raw.Wei(tc.bit)
			}

			if !reflect.DeepEqual(tc.wei, wei) {
				t.Fatalf("wei\n\n%s\n", cmp.Diff(tc.wei, wei))
			}
		})
	}
}
