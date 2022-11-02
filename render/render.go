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
		dot = append(dot, "edge [style=invis];")
		dot = append(dot, "graph [nodesep=1.0, rankdir=LR, style=dashed, ranksep=1.0];")
		dot = append(dot, "node [fontname=\"Courier\", shape=record];")
	}

	var rec func(getlin.Mapper, int)
	{
		rec = func(mpr getlin.Mapper, lev int) {
			for _, x := range mpr.All() {
				for j, y := range x {
					if y.Mapper() == nil {
						dot = append(dot, record(lev, y, mpr.Ind(y)))
						dot = append(dot, subgrp(lev, mpr.Ind(y)))
					} else {
						rec(y.Mapper(), lev+j+1)
					}

					for k, z := range mpr.Bel(y) {
						if z.Mapper() == nil {
							lin := linker(lev, mpr.Ind(y), lev, mpr.Ind(z))
							dot = append(dot, lin)
						} else {
							for l := range z.Mapper().All()[0] {
								lin := linker(lev, mpr.Ind(y), lev+k+1, l)
								dot = append(dot, lin)
							}
						}
					}
				}
			}
		}
	}

	{
		rec(mpr, 0)
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

func linker(pre int, lef int, nex int, rig int) string {
	return fmt.Sprintf("L%03dM%03d -> L%03dM%03d;", pre, lef, nex, rig)
}

func record(lev int, mod getlin.Module, ind int) string {
	return fmt.Sprintf(
		"L%03dM%03d [label=\"{%s}|{{Inp|%d}|{States|%d}|{Out|%d}}\"];",
		lev,
		ind,
		module(mod),
		mod.Shaper().Inp(),
		mod.Shaper().Sta(),
		mod.Shaper().Out(),
	)
}

func subgrp(lev int, ind int) string {
	return fmt.Sprintf("subgraph cluster%03d { L%03dM%03d }", lev, lev, ind)
}
