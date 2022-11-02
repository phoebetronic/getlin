package render

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
	"github.com/phoebetron/getlin/render"
)

func Test_Render_Create(t *testing.T) {
	testCases := []struct {
		mod [][]getlin.Module
	}{
		// Case 0
		{
			mod: [][]getlin.Module{
				{
					muslin(2, 2),
				},
			},
		},
		// Case 1
		{
			mod: [][]getlin.Module{
				{
					muslin(512, 32),
					muslin(512, 32),
				},
				{
					muslin(64, 16),
					muslin(64, 16),
					muslin(64, 16),
					muslin(64, 16),
				},
				{
					musgra([][]getlin.Module{
						{
							muslin(64, 16),
							muslin(64, 16),
						},
						{musvot(32)},
					}),
					musgra([][]getlin.Module{
						{muslin(64, 64)},
						{musvot(64)},
					}),
					musgra([][]getlin.Module{
						{muslin(64, 64)},
						{musvot(64)},
					}),
					musgra([][]getlin.Module{
						{muslin(64, 64)},
						{musvot(64)},
					}),
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var mpr getlin.Mapper
			{
				mpr = mapper.New(mapper.Config{Mod: tc.mod})
			}

			var act []byte
			{
				act = []byte(strings.Join(render.Create(mpr), " "))
			}

			var pat string
			{
				pat = filepath.Join("gold", fmt.Sprintf("%03d.gold.txt", i))
			}

			if *upd {
				err := os.WriteFile(pat, act, 0600)
				if err != nil {
					t.Fatal(err)
				}
			}

			var exp []byte
			{
				exp, err = os.ReadFile(pat)
				if err != nil {
					t.Fatal(err)
				}
			}

			if !bytes.Equal(exp, act) {
				t.Fatalf("byt\n\n%s\n", cmp.Diff(string(exp), string(act)))
			}
		})
	}
}
