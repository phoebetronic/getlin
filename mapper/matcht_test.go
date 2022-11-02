package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Mapper_matcht(t *testing.T) {
	testCases := []struct {
		lay []int
		out int
		ind int
		tru [2]int
	}{
		// Case 0
		{
			lay: []int{1},
			out: 1,
			ind: 0,
			tru: [2]int{0, 1},
		},

		// Case 1
		{
			lay: []int{2},
			out: 2,
			ind: 0,
			tru: [2]int{0, 2},
		},

		// Case 2
		{
			lay: []int{1, 1},
			out: 2,
			ind: 0,
			tru: [2]int{0, 1},
		},
		// Case 3
		{
			lay: []int{1, 1},
			out: 2,
			ind: 1,
			tru: [2]int{1, 2},
		},

		// Case 4
		{
			lay: []int{2, 2},
			out: 2,
			ind: 0,
			tru: [2]int{0, 1},
		},
		// Case 5
		{
			lay: []int{2, 2},
			out: 2,
			ind: 1,
			tru: [2]int{1, 2},
		},

		// Case 6
		{
			lay: []int{2, 3},
			out: 2,
			ind: 0,
			tru: [2]int{0, 0},
		},
		// Case 7
		{
			lay: []int{2, 3},
			out: 2,
			ind: 1,
			tru: [2]int{0, 0},
		},

		// Case 8
		{
			lay: []int{1, 1, 1, 1},
			out: 1,
			ind: 0,
			tru: [2]int{0, 1},
		},
		// Case 9
		{
			lay: []int{1, 1, 1, 1},
			out: 1,
			ind: 1,
			tru: [2]int{0, 1},
		},
		// Case 10
		{
			lay: []int{1, 1, 1, 1},
			out: 1,
			ind: 2,
			tru: [2]int{0, 1},
		},
		// Case 11
		{
			lay: []int{1, 1, 1, 1},
			out: 1,
			ind: 3,
			tru: [2]int{0, 1},
		},

		// Case 12
		{
			lay: []int{1, 1, 1, 1},
			out: 2,
			ind: 0,
			tru: [2]int{0, 1},
		},
		// Case 13
		{
			lay: []int{1, 1, 1, 1},
			out: 2,
			ind: 1,
			tru: [2]int{0, 1},
		},
		// Case 14
		{
			lay: []int{1, 1, 1, 1},
			out: 2,
			ind: 2,
			tru: [2]int{1, 2},
		},
		// Case 15
		{
			lay: []int{1, 1, 1, 1},
			out: 2,
			ind: 3,
			tru: [2]int{1, 2},
		},

		// Case 16
		{
			lay: []int{1, 2, 1},
			out: 2,
			ind: 0,
			tru: [2]int{0, 0},
		},
		// Case 17
		{
			lay: []int{1, 2, 1},
			out: 2,
			ind: 1,
			tru: [2]int{0, 0},
		},
		// Case 18
		{
			lay: []int{1, 2, 1},
			out: 2,
			ind: 2,
			tru: [2]int{0, 0},
		},

		// Case 19
		{
			lay: []int{1, 1, 1, 1},
			out: 4,
			ind: 0,
			tru: [2]int{0, 1},
		},
		// Case 20
		{
			lay: []int{1, 1, 1, 1},
			out: 4,
			ind: 1,
			tru: [2]int{1, 2},
		},
		// Case 21
		{
			lay: []int{1, 1, 1, 1},
			out: 4,
			ind: 2,
			tru: [2]int{2, 3},
		},
		// Case 22
		{
			lay: []int{1, 1, 1, 1},
			out: 4,
			ind: 3,
			tru: [2]int{3, 4},
		},

		// Case 23
		{
			lay: []int{4, 1, 1, 4},
			out: 10,
			ind: 2,
			tru: [2]int{5, 6},
		},

		// Case 24
		{
			lay: []int{3, 3, 3},
			out: 3,
			ind: 0,
			tru: [2]int{0, 1},
		},
		// Case 25
		{
			lay: []int{3, 3, 3},
			out: 3,
			ind: 1,
			tru: [2]int{1, 2},
		},
		// Case 26
		{
			lay: []int{3, 3, 3},
			out: 3,
			ind: 2,
			tru: [2]int{2, 3},
		},

		// Case 27
		{
			lay: []int{3, 3, 3},
			out: 15,
			ind: 0,
			tru: [2]int{0, 0},
		},
		// Case 28
		{
			lay: []int{3, 3, 3},
			out: 15,
			ind: 1,
			tru: [2]int{0, 0},
		},
		// Case 29
		{
			lay: []int{3, 3, 3},
			out: 15,
			ind: 2,
			tru: [2]int{0, 0},
		},

		// Case 30
		{
			lay: []int{16, 16, 16, 16},
			out: 8,
			ind: 0,
			tru: [2]int{0, 2},
		},
		// Case 31
		{
			lay: []int{16, 16, 16, 16},
			out: 8,
			ind: 1,
			tru: [2]int{2, 4},
		},
		// Case 32
		{
			lay: []int{16, 16, 16, 16},
			out: 8,
			ind: 2,
			tru: [2]int{4, 6},
		},
		// Case 33
		{
			lay: []int{16, 16, 16, 16},
			out: 8,
			ind: 3,
			tru: [2]int{6, 8},
		},

		// Case 34
		{
			lay: []int{2, 2},
			out: 4,
			ind: 0,
			tru: [2]int{0, 2},
		},
		// Case 35
		{
			lay: []int{2, 2},
			out: 4,
			ind: 1,
			tru: [2]int{2, 4},
		},

		// Case 36
		{
			lay: []int{2, 5, 9, 1},
			out: 1,
			ind: 0,
			tru: [2]int{0, 1},
		},
		// Case 37
		{
			lay: []int{2, 5, 9, 1},
			out: 1,
			ind: 1,
			tru: [2]int{0, 1},
		},
		// Case 38
		{
			lay: []int{2, 5, 9, 1},
			out: 1,
			ind: 2,
			tru: [2]int{0, 1},
		},
		// Case 39
		{
			lay: []int{2, 5, 9, 1},
			out: 1,
			ind: 3,
			tru: [2]int{0, 1},
		},

		// Case 40
		{
			lay: []int{5, 5, 5},
			out: 30,
			ind: 0,
			tru: [2]int{0, 0},
		},
		// Case 41
		{
			lay: []int{5, 5, 5},
			out: 30,
			ind: 1,
			tru: [2]int{0, 0},
		},
		// Case 42
		{
			lay: []int{5, 5, 5},
			out: 30,
			ind: 2,
			tru: [2]int{0, 0},
		},

		// Case 43
		{
			lay: []int{3, 3, 3},
			out: 9,
			ind: 0,
			tru: [2]int{0, 3},
		},
		// Case 44
		{
			lay: []int{3, 3, 3},
			out: 9,
			ind: 1,
			tru: [2]int{3, 6},
		},
		// Case 45
		{
			lay: []int{3, 3, 3},
			out: 9,
			ind: 2,
			tru: [2]int{6, 9},
		},

		// Case 46
		{
			lay: []int{2, 2, 2},
			out: 2,
			ind: 0,
			tru: [2]int{0, 0},
		},
		// Case 47
		{
			lay: []int{2, 2, 2},
			out: 2,
			ind: 1,
			tru: [2]int{0, 0},
		},
		// Case 48
		{
			lay: []int{2, 2, 2},
			out: 2,
			ind: 2,
			tru: [2]int{0, 0},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var tru [2]int
			{
				tru = matcht(tc.lay, tc.out, tc.ind)
			}

			if !reflect.DeepEqual(tru, tc.tru) {
				t.Fatalf("tru\n\n%s\n", cmp.Diff(tc.tru, tru))
			}
		})
	}
}
