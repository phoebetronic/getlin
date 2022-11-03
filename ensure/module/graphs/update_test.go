package graphs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
	"github.com/phoebetron/getlin/module/graphs"
)

func Test_Module_Graphs_Update_Metric(t *testing.T) {
	testCases := []struct {
		inp getlin.Vector
		mod getlin.Module
		bef [41]float32
		aft [41]float32
	}{
		// Case 0
		{
			inp: musvec([]int{0, 1}, []int{0, 1}),
			mod: musgra([][]getlin.Module{
				{
					muslin(2, 1),
					muslin(2, 1),
				},
				{
					muslin(2, 2),
				},
			}),
			bef: [41]float32{},
			aft: [41]float32{
				0.1875,
				0,
				0,
				0,
				0,
				0.0625,
				0,
				0,
				0,
				0,
				0.125,
				0,
				0,
				0,
				0,
				0.125,
				0,
				0,
				0,
				0, // -
				0, // 0
				0, // +
				0,
				0,
				0,
				0.125,
				0,
				0,
				0,
				0,
				0.125,
				0,
				0,
				0,
				0,
				0.1875,
				0,
				0,
				0,
				0,
				0.0625,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var bef [41]float32
			{
				bef = tc.mod.Metric().Get().Sta().Nrm()
			}

			if !reflect.DeepEqual(bef, tc.bef) {
				t.Fatalf("bef\n\n%s\n", cmp.Diff(tc.bef, bef))
			}

			{
				tc.mod.Update(tc.inp)
			}

			var aft [41]float32
			{
				aft = tc.mod.Metric().Get().Sta().Nrm()
			}

			if !reflect.DeepEqual(aft, tc.aft) {
				t.Fatalf("aft\n\n%s\n", cmp.Diff(tc.aft, aft))
			}
		})
	}
}

func Test_Module_Graphs_Update_States(t *testing.T) {
	testCases := []struct {
		inp getlin.Vector
		mod getlin.Module
		bef []float32
		aft []float32
	}{
		// Case 0
		{
			inp: musvec([]int{0, 1}, []int{0, 1}),
			mod: musgra([][]getlin.Module{
				{
					muslin(2, 1),
					muslin(2, 1),
				},
				{
					muslin(2, 2),
				},
			}),
			bef: []float32{
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
				0,
			},
			aft: []float32{
				+4,
				+2,
				-2,
				-3,
				-4,
				-1,
				+3,
				+3,
				+4,
				+4,
				-4,
				-1,
				-3,
				-2,
				+4,
				+1,
			},
		},
		// Case 1
		{
			inp: musvec([]int{0, 1, 0}, []int{0, 1}),
			mod: musgra([][]getlin.Module{
				{
					muslin(3, 2),
					muslin(3, 2),
					muslin(3, 2),
					muslin(3, 2),
				},
				{
					musgra([][]getlin.Module{
						{
							muslin(8, 2),
							muslin(8, 2),
						},
						{
							musvot(4),
						},
					}),
					musgra([][]getlin.Module{
						{
							muslin(8, 2),
							muslin(8, 2),
						},
						{
							musvot(4),
						},
					}),
				},
			}),
			bef: []float32{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			aft: []float32{
				+4, +2, +2, -4, -1, -4, +2, +2,
				+3, -3, -3, -4, +2, +3, +4, -4,
				-1, -2, +2, +2, +1, -2, -4, -2,
				-1, -4, -4, +3, +4, +3, -2, -3,
				-2, +1, +1, +3, -4, -4, -3, +3,
				+1, +3, -4, -2, -1, +3, +4, +1,
				+4, +2, +2, +2, +4, +3, +3, +1,
				-2, -2, -4, -4, -3, -1, -4, -4,
				+1, +2, +2, +4, +4, +3, +3, +3,
				-4, -3, -4, -2, -3, -2, -1, -2,
				+4, +1, +4, +4, +3, +2, +3, +2,
				-3, -2, -2, -3, -3, -4, -3, -3,
				+4, +1, +2, +4, +1, +1, +1, +1,
				-2, -4, -3, -2, -3, -4, -4, -1,
				-1, -1, -4, -1, -1, -2, -2, -3,
				+4, +3, +1, +4, +1, +4, +3, +3,
				-4, -3, -3, -3, -4, -3, -4, -1,
				+4, +3, +4, +1, +1, +3, +1, +3,
				-2, -4, -3, -1, -1, -3, -4, -2,
				+4, +4, +2, +4, +4, +1, +3, +3,
				-2, -1, -4, -1, -2, -4, -3, -4,
				+1, +1, +3, +1, +1, +2, +1, +3,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var bef []float32
			{
				bef = tc.mod.States()
			}

			if !reflect.DeepEqual(bef, tc.bef) {
				t.Fatalf("bef\n\n%s\n", cmp.Diff(tc.bef, bef))
			}

			{
				tc.mod.Update(tc.inp)
			}

			var aft []float32
			{
				aft = tc.mod.States()
			}

			if !reflect.DeepEqual(aft, tc.aft) {
				t.Fatalf("aft\n\n%s\n", cmp.Diff(tc.aft, aft))
			}
		})
	}
}

func Test_Module_Graphs_Update_Labels_Linear(t *testing.T) {
	var mo0 *tester
	var mo1 *tester
	var mo2 getlin.Module
	var mo3 *tester
	var mo4 *tester
	var mo5 *tester
	var mo6 getlin.Module
	var mo7 *tester
	var mo8 *tester
	var mo9 *tester
	{
		mo0 = &tester{Sea: musvec(nil, []int{0, 0, 0, 1}), Sha: mussha(3, 4)}
		mo1 = &tester{Sea: musvec(nil, []int{0, 0, 1, 0}), Sha: mussha(3, 4)}

		mo3 = &tester{Sea: musvec(nil, []int{0, 0}), Sha: mussha(8, 2)}
		mo4 = &tester{Sea: musvec(nil, []int{1, 1}), Sha: mussha(8, 2)}
		mo5 = &tester{Sea: musvec(nil, []int{0, 1, 0, 1}), Sha: mussha(4, 4)}
		mo2 = musgra([][]getlin.Module{
			{
				mo3,
				mo4,
			},
			{
				mo5,
			},
		})

		mo7 = &tester{Sea: musvec(nil, []int{1, 0}), Sha: mussha(8, 2)}
		mo8 = &tester{Sea: musvec(nil, []int{0, 1}), Sha: mussha(8, 2)}
		mo9 = &tester{Sea: musvec(nil, []int{0, 0, 1, 1}), Sha: mussha(4, 4)}
		mo6 = musgra([][]getlin.Module{
			{
				mo7,
				mo8,
			},
			{
				mo9,
			},
		})
	}

	testCases := []struct {
		vec getlin.Vector
		mod *tester
		upd getlin.Vector
	}{
		// Case 0
		{
			vec: musvec([]int{1, 1, 1}, []int{0, 1, 0, 1, 0, 0, 1, 1}),
			mod: mo1,
			upd: musvec([]int{1, 1, 1}, []int{0, 0, 1, 0}),
		},
		// Case 1
		{
			vec: musvec([]int{1, 1, 1}, []int{0, 0, 0, 0, 1, 0, 1, 0}),
			mod: mo1,
			upd: musvec([]int{1, 1, 1}, []int{1, 0, 1, 0}),
		},
		// Case 2
		{
			vec: musvec([]int{1, 1, 1}, []int{0, 1, 0, 1, 0, 0, 1, 1}),
			mod: mo5,
			upd: musvec([]int{0, 0, 1, 1}, []int{0, 1, 0, 1}),
		},
		// Case 3
		{
			vec: musvec([]int{1, 1, 1}, []int{1, 1, 1, 1, 0, 1, 1, 0}),
			mod: mo5,
			upd: musvec([]int{0, 0, 1, 1}, []int{1, 1, 1, 1}),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var mod getlin.Module
			{
				mod = graphs.New(graphs.Config{
					Mpr: mapper.New(mapper.Config{
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo6,
							},
						},
					}),
				})
			}

			{
				mod.Update(tc.vec)
			}

			if !reflect.DeepEqual(tc.mod.Upd.Inp().Raw(), tc.upd.Inp().Raw()) {
				t.Fatalf("inp\n\n%s\n", cmp.Diff(tc.upd.Inp().Raw(), tc.mod.Upd.Inp().Raw()))
			}
			if !reflect.DeepEqual(tc.mod.Upd.Out().Raw(), tc.upd.Out().Raw()) {
				t.Fatalf("out\n\n%s\n", cmp.Diff(tc.upd.Out().Raw(), tc.mod.Upd.Out().Raw()))
			}
		})
	}
}

func Test_Module_Graphs_Update_Labels_Voting(t *testing.T) {
	var mo0 *tester
	var mo1 *tester
	var mo2 getlin.Module
	var mo3 *tester
	var mo4 *tester
	var mo5 *tester
	var mo6 getlin.Module
	var mo7 *tester
	var mo8 *tester
	var mo9 *tester
	{
		mo0 = &tester{Sea: musvec(nil, []int{0, 0, 1}), Sha: mussha(3, 3)}
		mo1 = &tester{Sea: musvec(nil, []int{0, 1, 0}), Sha: mussha(3, 3)}

		mo3 = &tester{Sea: musvec(nil, []int{0, 0, 0, 1}), Sha: mussha(6, 4)}
		mo4 = &tester{Sea: musvec(nil, []int{0, 0, 1, 0}), Sha: mussha(6, 4)}
		mo5 = &tester{Sea: musvec(nil, []int{1}), Sha: mussha(8, 1)}
		mo2 = musgra([][]getlin.Module{
			{
				mo3,
				mo4,
			},
			{
				mo5,
			},
		})

		mo7 = &tester{Sea: musvec(nil, []int{0, 0, 1, 1}), Sha: mussha(6, 4)}
		mo8 = &tester{Sea: musvec(nil, []int{0, 1, 0, 0}), Sha: mussha(6, 4)}
		mo9 = &tester{Sea: musvec(nil, []int{0}), Sha: mussha(8, 1)}
		mo6 = musgra([][]getlin.Module{
			{
				mo7,
				mo8,
			},
			{
				mo9,
			},
		})
	}

	testCases := []struct {
		vec getlin.Vector
		mod *tester
		upd getlin.Vector
	}{
		// Case 0, the true labels are defined as [1 0], which is what the
		// configured system actually predicts, see mo5 and mo9. Here mo0
		// predicts [0 0 1], which, upon successful prediction, must be the true
		// labels received for mo0 in Module.Update.
		{
			vec: musvec([]int{1, 1, 1}, []int{1, 0}),
			mod: mo0,
			upd: musvec([]int{1, 1, 1}, []int{0, 0, 1}),
		},
		// Case 1, the true labels are defined as [1 1], which is NOT what the
		// configured system actually predicts, see mo5 and mo9. Here mo0 must
		// receive the scaled version of the true labels defined by the original
		// input Vector. That is, based on the true label index range [0 1], [1
		// 1 1].
		{
			vec: musvec([]int{1, 1, 1}, []int{1, 1}),
			mod: mo0,
			upd: musvec([]int{1, 1, 1}, []int{1, 1, 1}),
		},
		// Case 2, the true labels are defined as [0 1], which is NOT what the
		// configured system actually predicts, see mo5 and mo9. Here mo0 must
		// receive the scaled version of the true labels defined by the original
		// input Vector. That is, based on the true label index range [0 1], [0
		// 0 0].
		{
			vec: musvec([]int{1, 1, 1}, []int{0, 1}),
			mod: mo0,
			upd: musvec([]int{1, 1, 1}, []int{0, 0, 0}),
		},
		// Case 3, the true labels are defined as [1 0], which is what the
		// configured system actually predicts, see mo5 and mo9. Here mo8
		// predicts [0 1 0 0], which, upon successful prediction, must be the
		// true labels received for mo8 in Module.Update.
		{
			vec: musvec([]int{1, 1, 1}, []int{1, 0}),
			mod: mo8,
			upd: musvec([]int{0, 0, 1, 0, 1, 0}, []int{0, 1, 0, 0}),
		},
		// Case 4, the true labels are defined as [1 1], which is NOT what the
		// configured system actually predicts, see mo5 and mo9. Here mo8 must
		// receive the scaled version of the true labels defined by the original
		// input Vector. That is, based on the true label index range [1 2], [1
		// 1 1 1].
		{
			vec: musvec([]int{1, 1, 1}, []int{1, 1}),
			mod: mo8,
			upd: musvec([]int{0, 0, 1, 0, 1, 0}, []int{1, 1, 1, 1}),
		},
		// Case 5, the true labels are defined as [0 0], which is NOT what the
		// configured system actually predicts, see mo5 and mo9. Here mo8 must
		// receive the scaled version of the true labels defined by the original
		// input Vector. That is, based on the true label index range [0 1], [0
		// 0 0 0].
		{
			vec: musvec([]int{1, 1, 1}, []int{0, 0}),
			mod: mo8,
			upd: musvec([]int{0, 0, 1, 0, 1, 0}, []int{0, 0, 0, 0}),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var mod getlin.Module
			{
				mod = graphs.New(graphs.Config{
					Mpr: mapper.New(mapper.Config{
						Mod: [][]getlin.Module{
							{
								mo0,
								mo1,
							},
							{
								mo2,
								mo6,
							},
						},
					}),
				})
			}

			{
				mod.Update(tc.vec)
			}

			if !reflect.DeepEqual(tc.mod.Upd.Inp().Raw(), tc.upd.Inp().Raw()) {
				t.Fatalf("inp\n\n%s\n", cmp.Diff(tc.upd.Inp().Raw(), tc.mod.Upd.Inp().Raw()))
			}
			if !reflect.DeepEqual(tc.mod.Upd.Out().Raw(), tc.upd.Out().Raw()) {
				t.Fatalf("out\n\n%s\n", cmp.Diff(tc.upd.Out().Raw(), tc.mod.Upd.Out().Raw()))
			}
		})
	}
}
