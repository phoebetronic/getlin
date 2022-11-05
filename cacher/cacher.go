package cacher

import (
	"fmt"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

type Cacher struct {
	ind int
	out map[int][]getlin.Binary
	vec map[int]getlin.Vector
}

func New(con Config) *Cacher {
	{
		con.Verify()
	}

	return &Cacher{
		ind: -1,
		out: map[int][]getlin.Binary{},
		vec: map[int]getlin.Vector{},
	}
}

func (c *Cacher) Add(inp getlin.Binary) {
	c.ind++
	c.vec[c.ind] = vector.New(vector.Config{Inp: inp.Raw()})
}

func (c *Cacher) Del() {
	c.ind = -1
	c.out = map[int][]getlin.Binary{}
	c.vec = map[int]getlin.Vector{}
}

func (c *Cacher) Lat() getlin.Vector {
	x := c.vec[c.ind]
	return x
}

func (c *Cacher) Log() {
	fmt.Printf("%4s %15s %15s\n", "lay", "inp", "out")
	fmt.Printf("\n")

	for k, v := range c.vec {
		fmt.Printf(
			"%4d %15s %15s\n",
			k,
			fmt.Sprintf("%v", v.Inp().Raw()),
			fmt.Sprintf("%v", v.Out().Raw()),
		)
	}
}

func (c *Cacher) Out(ind int) []getlin.Binary {
	return c.out[ind]
}

func (c *Cacher) Upd(out getlin.Binary) {
	var vec getlin.Vector
	{
		vec = c.vec[c.ind]
	}

	{
		c.out[c.ind] = append(c.out[c.ind], out)
	}

	for _, x := range out.Raw() {
		vec.Out().Add(x)
	}

	{
		c.vec[c.ind] = vec
	}
}

func (c *Cacher) Vec(ind int) getlin.Vector {
	return c.vec[ind]
}
