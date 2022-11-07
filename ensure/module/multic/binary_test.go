package multic

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/loader/binary"
	"github.com/phoebetron/getlin/module/multic"
	"github.com/phoebetron/getlin/module/single"
	"github.com/phoebetron/getlin/random/native"
)

func Test_Module_Multic_Binary(t *testing.T) {
	{
		rand.Seed(1)
	}

	var ldr getlin.Loader
	{
		ldr = binary.New(binary.Config{})
	}

	var con single.Config
	{
		con = single.Config{
			Cla: 64,
			Fre: 0.1,
			Inp: 4,
			Ran: native.New(native.Config{}),
			Sta: 32,
			Thr: 4,
		}
	}

	var mod getlin.Module
	{
		mod = multic.New(multic.Config{
			Sin: []getlin.Module{
				single.New(con),
				single.New(con),
			},
		})
	}

	var met getlin.Metric
	for i := 1; i <= 5; i++ {
		var bat [][2]getlin.Vector
		{
			bat = ldr.Search()
		}

		for _, x := range bat {
			mod.Update(x[0])
			mod.Update(x[1])
		}

		{
			met = mod.Verify(bat)
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

	var bat [][2]getlin.Vector
	{
		bat = ldr.Search()
	}

	{
		fmt.Printf(
			"The test data defines %b in class 1 to be %b, which the Module confirms with %b.\n",
			bat[9][0].Inp().Raw(),
			bat[9][0].Out().Raw(),
			mod.Search(bat[9][0]).Out().Raw(),
		)
	}

	{
		fmt.Printf(
			"The test data defines %b in class 1 to be %b, which the Module confirms with %b.\n",
			bat[9][1].Inp().Raw(),
			bat[9][1].Out().Raw(),
			mod.Search(bat[9][1]).Out().Raw(),
		)
	}

	{
		fmt.Printf("\n")
	}

	{
		fmt.Printf(
			"The test data defines %b in class 2 to be %b, which the Module confirms with %b.\n",
			bat[21][0].Inp().Raw(),
			bat[21][0].Out().Raw(),
			mod.Search(bat[21][0]).Out().Raw(),
		)
	}

	{
		fmt.Printf(
			"The test data defines %b in class 2 to be %b, which the Module confirms with %b.\n",
			bat[21][1].Inp().Raw(),
			bat[21][1].Out().Raw(),
			mod.Search(bat[21][1]).Out().Raw(),
		)
	}

	{
		fmt.Printf("\n")
	}
}
