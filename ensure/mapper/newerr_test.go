package mapper

import (
	"fmt"
	"testing"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
)

func Test_Mapper_New_Err(t *testing.T) {
	testCases := []struct {
		mod [][]getlin.Module
	}{
		// Case 0
		{
			mod: [][]getlin.Module{
				{
					muslin(2, 2),
					muslin(2, 2),
				},
				{
					muslin(3, 2), // 3 inputs do not match 4 outputs
					muslin(4, 2),
				},
			},
		},
		// Case 1
		{
			mod: [][]getlin.Module{
				{
					muslin(2, 2),
					muslin(2, 2),
				},
				{
					muslin(4, 2), // 5 outputs do not match 4 outputs
					muslin(4, 3),
				},
			},
		},
		// Case 2
		{
			mod: [][]getlin.Module{
				{
					muslin(3, 3),
					muslin(7, 3),
				},
				{
					muslin(6, 1),
					muslin(6, 1),
					muslin(6, 1),
				},
				{
					muslin(3, 7),
					muslin(3, 3),
				},
			},
		},
		// Case 3
		{
			mod: [][]getlin.Module{
				{
					muslin(3, 3),
					muslin(7, 3),
				},
				{
					muslin(6, 1),
					muslin(6, 1),
					muslin(6, 6),
				},
				{
					muslin(8, 6),
					muslin(8, 2),
				},
			},
		},
		// Case 4
		{
			mod: [][]getlin.Module{
				{
					muslin(2, 3),
					muslin(6, 2),
					muslin(2, 3), // 3 modules here cannot receive feedback from 2 modules below
				},
				{
					musvot(4),
					musvot(4),
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var erc chan error
			{
				erc = make(chan error, 100)
			}

			{
				mapper.New(mapper.Config{
					Err: func(err error) { erc <- err },
					Mod: tc.mod,
				})
			}

			{
				close(erc)
			}

			if !mapper.IsInvalidMapping(<-erc) {
				t.Fatal("expected error to be invalidMappingError")
			}
		})
	}
}
