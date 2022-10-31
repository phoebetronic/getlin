package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/module/linear"
	"github.com/phoebetron/getlin/module/static"
	"github.com/phoebetron/getlin/vector"
)

func Test_Mapper(t *testing.T) {
	var mo1 getlin.Module
	var mo2 getlin.Module
	var mo3 getlin.Module
	var mo4 getlin.Module
	var mo5 getlin.Module
	var mo6 getlin.Module
	var mo7 getlin.Module
	{
		mo1 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo2 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo3 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo4 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo5 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo6 = static.New(static.Config{Out: vector.New(vector.Config{})})
		mo7 = static.New(static.Config{Out: vector.New(vector.Config{})})
	}

	testCases := []struct {
		mod getlin.Module
		ind int
		lay int
		abo []getlin.Module
		bel []getlin.Module
	}{
		// Case 0
		{
			mod: mo1,
			ind: 0,
			lay: 0,
			abo: nil,
			bel: []getlin.Module{
				mo3,
				mo4,
				mo5,
				mo6,
			},
		},
		// Case 1
		{
			mod: mo2,
			ind: 1,
			lay: 0,
			abo: nil,
			bel: []getlin.Module{
				mo3,
				mo4,
				mo5,
				mo6,
			},
		},
		// Case 2
		{
			mod: mo3,
			ind: 2,
			lay: 1,
			abo: []getlin.Module{
				mo1,
				mo2,
			},
			bel: []getlin.Module{
				mo7,
			},
		},
		// Case 3
		{
			mod: mo5,
			ind: 4,
			lay: 1,
			abo: []getlin.Module{
				mo1,
				mo2,
			},
			bel: []getlin.Module{
				mo7,
			},
		},
		// Case 4
		{
			mod: mo7,
			ind: 6,
			lay: 2,
			abo: []getlin.Module{
				mo3,
				mo4,
				mo5,
				mo6,
			},
			bel: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var mpr getlin.Mapper
			{
				mpr = New(Config{
					Mod: [][]getlin.Module{
						{
							mo1,
							mo2,
						},
						{
							mo3,
							mo4,
							mo5,
							mo6,
						},
						{
							mo7,
						},
					},
				})
			}

			var abo []getlin.Module
			{
				abo = mpr.Abo(tc.mod)
			}

			var bel []getlin.Module
			{
				bel = mpr.Bel(tc.mod)
			}

			var ind int
			{
				ind = mpr.Ind(tc.mod)
			}

			var lay int
			{
				lay = mpr.Lay(tc.mod)
			}

			if !reflect.DeepEqual(abo, tc.abo) {
				t.Fatalf("modules above do not match")
			}
			if !reflect.DeepEqual(bel, tc.bel) {
				t.Fatalf("modules below do not match")
			}
			if !reflect.DeepEqual(ind, tc.ind) {
				t.Fatalf("ind\n\n%s\n", cmp.Diff(tc.ind, ind))
			}
			if !reflect.DeepEqual(lay, tc.lay) {
				t.Fatalf("lay\n\n%s\n", cmp.Diff(tc.lay, lay))
			}
		})
	}
}

func Test_Mapper_Tru(t *testing.T) {
	var erc chan error

	testCases := []struct {
		set func() (getlin.Mapper, getlin.Module)
		tru []int
	}{
		// Case 0
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				{
					mo0 = muslin(2, 1)
					mo1 = muslin(2, 1)
					mo2 = muslin(2, 2)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
							},
						},
					})
				}

				return mpr, mo1
			},
			tru: []int{1, 2},
		},
		// Case 1
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				{
					mo0 = muslin(2, 1)
					mo1 = muslin(2, 1)
					mo2 = muslin(2, 2)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
							},
						},
					})
				}

				return mpr, mo2
			},
			tru: []int{0, 2},
		},
		// Case 2
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				{
					mo0 = muslin(2, 1)
					mo1 = muslin(2, 1)
					mo2 = muslin(2, 2)
					mo3 = muslin(2, 1)
					mo4 = muslin(2, 1)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
							},
							{
								mo3,
								mo4,
							},
						},
					})
				}

				return mpr, mo1
			},
			tru: []int{1, 2},
		},
		// Case 3
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				{
					mo0 = muslin(2, 1)
					mo1 = muslin(2, 1)
					mo2 = muslin(2, 2)
					mo3 = muslin(2, 1)
					mo4 = muslin(2, 1)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
							},
							{
								mo3,
								mo4,
							},
						},
					})
				}

				return mpr, mo0
			},
			tru: []int{0, 1},
		},
		// Case 4
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				{
					mo0 = muslin(2, 1)
					mo1 = muslin(2, 1)
					mo2 = muslin(2, 2)
					mo3 = static.New(static.Config{Out: vector.New(vector.Config{})})
					mo4 = muslin(2, 1)
					mo5 = muslin(2, 1)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
							},
							{
								mo4,
								mo5,
							},
						},
					})
				}

				return mpr, mo3
			},
			tru: nil, // static is purely functional and has no true labels
		},
		// Case 5
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				{
					mo0 = muslin(2, 1)
					mo1 = muslin(2, 1)
					mo2 = muslin(2, 2)
					mo3 = static.New(static.Config{Out: vector.New(vector.Config{})})
					mo4 = muslin(2, 1)
					mo5 = muslin(2, 1)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
							},
							{
								mo4,
								mo5,
							},
						},
					})
				}

				return mpr, mo4
			},
			tru: []int{0, 1},
		},
		// Case 6
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				var mo6 getlin.Module
				var mo7 getlin.Module
				{
					mo0 = muslin(3, 2)
					mo1 = muslin(7, 6)
					mo2 = muslin(8, 2)
					mo3 = muslin(8, 2)
					mo4 = muslin(8, 2)
					mo5 = muslin(8, 2)
					mo6 = muslin(8, 6)
					mo7 = muslin(8, 2)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
								mo4,
								mo5,
							},
							{
								mo6,
								mo7,
							},
						},
					})
				}

				return mpr, mo1
			},
			tru: []int{2, 8},
		},
		// Case 7
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				var mo6 getlin.Module
				var mo7 getlin.Module
				{
					mo0 = muslin(3, 2)
					mo1 = muslin(7, 6)
					mo2 = muslin(8, 2)
					mo3 = muslin(8, 2)
					mo4 = muslin(8, 2)
					mo5 = muslin(8, 2)
					mo6 = muslin(8, 6)
					mo7 = muslin(8, 2)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
								mo4,
								mo5,
							},
							{
								mo6,
								mo7,
							},
						},
					})
				}

				return mpr, mo2
			},
			tru: []int{0, 2},
		},
		// Case 8
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				var mo6 getlin.Module
				var mo7 getlin.Module
				{
					mo0 = muslin(3, 2)
					mo1 = muslin(7, 6)
					mo2 = muslin(8, 2)
					mo3 = muslin(8, 2)
					mo4 = muslin(8, 2)
					mo5 = muslin(8, 2)
					mo6 = muslin(8, 6)
					mo7 = muslin(8, 2)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
								mo4,
								mo5,
							},
							{
								mo6,
								mo7,
							},
						},
					})
				}

				return mpr, mo3
			},
			tru: []int{2, 4},
		},
		// Case 9
		{
			set: func() (getlin.Mapper, getlin.Module) {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				var mo6 getlin.Module
				{
					mo0 = muslin(2, 3)
					mo1 = muslin(5, 4)
					mo2 = muslin(7, 2)
					mo3 = muslin(7, 2)
					mo4 = muslin(7, 3)
					mo5 = muslin(7, 5)
					mo6 = muslin(7, 2)
				}

				var mpr getlin.Mapper
				{
					mpr = New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
								mo4,
							},
							{
								mo5,
								mo6,
							},
						},
					})
				}

				return mpr, mo4
			},
			tru: []int{4, 7},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			{
				erc = make(chan error, 100)
			}

			var mpr getlin.Mapper
			var mod getlin.Module
			{
				mpr, mod = tc.set()
			}

			{
				close(erc)
			}

			var err error
			{
				err = <-erc
			}

			if err != nil {
				t.Fatal(err)
			}

			var tru []int
			{
				tru = mpr.Tru(mod)
			}

			if !reflect.DeepEqual(tru, tc.tru) {
				t.Fatalf("tru\n\n%s\n", cmp.Diff(tc.tru, tru))
			}
		})
	}
}

func Test_Mapper_Tru_Err(t *testing.T) {
	var erc chan error

	testCases := []struct {
		set func()
	}{
		// Case 0
		{
			set: func() {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				{
					mo0 = muslin(2, 2)
					mo1 = muslin(2, 2)
					mo2 = muslin(3, 2) // 3 inputs do not match 4 outputs
					mo3 = muslin(4, 2)
				}

				{
					New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
							},
						},
					})
				}
			},
		},
		// Case 1
		{
			set: func() {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				{
					mo0 = muslin(2, 2)
					mo1 = muslin(2, 2)
					mo2 = muslin(4, 2)
					mo3 = muslin(4, 3) // 5 outputs do not match 4 outputs
				}

				{
					New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
							},
						},
					})
				}
			},
		},
		// Case 2
		{
			set: func() {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				var mo6 getlin.Module
				{
					mo0 = muslin(3, 3)
					mo1 = muslin(7, 3)
					mo2 = muslin(6, 1)
					mo3 = muslin(6, 1)
					mo4 = muslin(6, 1)
					mo5 = muslin(3, 7)
					mo6 = muslin(3, 3)
				}

				{
					New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
								mo4,
							},
							{
								mo5,
								mo6,
							},
						},
					})
				}
			},
		},
		// Case 3
		{
			set: func() {
				var mo0 getlin.Module
				var mo1 getlin.Module
				var mo2 getlin.Module
				var mo3 getlin.Module
				var mo4 getlin.Module
				var mo5 getlin.Module
				var mo6 getlin.Module
				{
					mo0 = muslin(3, 3)
					mo1 = muslin(7, 3)
					mo2 = muslin(6, 1)
					mo3 = muslin(6, 1)
					mo4 = muslin(6, 6)
					mo5 = muslin(8, 6)
					mo6 = muslin(8, 2)
				}

				{
					New(Config{
						Err: func(err error) { erc <- err },
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo3,
								mo4,
							},
							{
								mo5,
								mo6,
							},
						},
					})
				}
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			{
				erc = make(chan error, 100)
			}

			{
				tc.set()
			}

			{
				close(erc)
			}

			if !IsInvalidMapping(<-erc) {
				t.Fatal("expected error to be invalidMappingError")
			}
		})
	}
}

func muslin(inp int, out int) getlin.Module {
	return linear.New(linear.Config{Inp: inp, Out: out, Sta: 4})
}
