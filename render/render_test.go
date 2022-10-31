package render

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
	"github.com/phoebetron/getlin/module/linear"
	"github.com/phoebetron/getlin/module/static"
	"github.com/phoebetron/getlin/vector"
)

var upd = flag.Bool("u", false, "whether to update golden files")

func Test_Render_Create(t *testing.T) {
	testCases := []struct {
		mod [][]getlin.Module
	}{
		// Case 0
		{
			mod: [][]getlin.Module{
				{
					mussta(),
				},
			},
		},
		// Case 1
		{
			mod: [][]getlin.Module{
				{
					mussta(),
					mussta(),
					mussta(),
				},
				{
					mussta(),
					mussta(),
				},
			},
		},
		// Case 2
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
					muslin(64, 64),
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
				act = []byte(strings.Join(Create(mpr), " "))
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

func muslin(inp int, out int) getlin.Module {
	return linear.New(linear.Config{Inp: inp, Out: out, Sta: 4})
}

func mussta() getlin.Module {
	return static.New(static.Config{Out: vector.New(vector.Config{})})
}
