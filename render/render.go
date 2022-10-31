package render

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/phoebetron/getlin"
)

func Create(mpr getlin.Mapper) []string {
	var dot []string
	{
		dot = append(dot, "digraph {")
		dot = append(dot, "rankdir=LR;")
	}

	for _, x := range mpr.All() {
		for _, y := range x {
			{
				dot = append(dot, record(y, mpr.Ind(y)))
			}

			for _, z := range mpr.Bel(y) {
				dot = append(dot, linker(mpr.Ind(y), mpr.Ind(z)))
			}
		}
	}

	{
		dot = append(dot, "}")
	}

	return dot
}

// Render visualizes the Module architecture in SVG format, which may only be
// useful for a Graphs Module implementation.
func Render(dot []string) []byte {
	var err error

	var cmd *exec.Cmd
	{
		cmd = exec.Command("dot", "-Tsvg")
	}

	{
		cmd.Stdin = strings.NewReader(strings.Join(dot, " "))
	}

	var byt []byte
	{
		byt, err = cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
	}

	return byt
}

func module(mod getlin.Module) string {
	split := strings.Split(fmt.Sprintf("%T", mod), ".")

	if len(split) < 2 {
		panic("cannot parse module name")
	}

	return strings.Replace(split[0], "*", "", 1)
}

func linker(lef int, rig int) string {
	return fmt.Sprintf("M%03d -> M%03d;", lef, rig)
}

func record(mod getlin.Module, ind int) string {
	var sta int
	if mod.Clause() != 0 {
		sta = mod.Automa() / 2 / mod.Clause()
	}

	return fmt.Sprintf(
		"M%03d [shape=record, label=\"{%s}|{{Inp|%d}|{States|%d}|{Out|%d}}\", fontname=\"Courier\"];",
		ind,
		module(mod),
		sta,
		mod.Automa(),
		mod.Clause(),
	)
}
