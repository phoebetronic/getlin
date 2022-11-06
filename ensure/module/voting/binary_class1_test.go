package render

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/loader/binary"
	"github.com/phoebetron/getlin/module/voting"
	"github.com/phoebetron/getlin/random/native"
)

func Test_Module_Voting_Binary_Class1(t *testing.T) {
	{
		rand.Seed(1)
	}

	var cl1 = 1

	var ldr getlin.Loader
	{
		ldr = binary.New(binary.Config{})
	}

	var mod getlin.Module
	{
		mod = voting.New(voting.Config{
			Cla: 64,
			Fre: 0.1,
			Inp: 4,
			Ran: native.New(native.Config{}),
			Sta: 32,
			Thr: 4,
		})
	}

	var met getlin.Metric
	for i := 1; i <= 5; i++ {
		var bat map[int][][2]getlin.Vector
		{
			bat = ldr.Search()
		}

		for _, x := range bat[cl1] {
			mod.Update(x[0])
			mod.Update(x[1])
		}

		{
			met = mod.Verify(bat[cl1])
		}

		fmt.Printf(
			"epo %2d    mae %4.3f    acc %4.3f\n",
			i,
			met.Get().Mat().Mae(),
			met.Get().Mat().Acc(),
		)
	}

	if met.Get().Mat().Acc() != 1.0 {
		t.Fatal("accuracy must be 1.0")
	}

	{
		fmt.Printf("\n")
	}

	var bat map[int][][2]getlin.Vector
	{
		bat = ldr.Search()
	}

	{
		fmt.Printf(
			"The test data defines %v to be %b, which the Module confirms with %b.\n",
			bat[cl1][6][0].Inp().Raw(),
			bat[cl1][6][0].Out().Raw()[0],
			mod.Search(bat[cl1][6][0]).Out().Raw()[0],
		)
	}

	{
		fmt.Printf(
			"The test data defines %v to be %b, which the Module confirms with %b.\n",
			bat[cl1][6][1].Inp().Raw(),
			bat[cl1][6][1].Out().Raw()[0],
			mod.Search(bat[cl1][6][1]).Out().Raw()[0],
		)
	}

	{
		fmt.Printf("\n")
	}
}
