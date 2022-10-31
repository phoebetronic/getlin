# harrel

A form of hard ReLu activation. The minimum threshold may not be lower than 0.05
in any case. That ensures for instance neutral automas having their internal
pointer states exactly at zero to be updated at least 5% of the time. Above that
minimum threshold, activation probability grows linearly like a classical ReLu.
The following figure can be rendered by running the code below via `go run
main.go`.

### figure

![Render](/assets/harrel.svg)

### code

```go
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/active/harrel"
	"github.com/phoebetron/getlin/random/crypto"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	var act getlin.Active
	{
		act = harrel.New(harrel.Config{})
	}

	var ran getlin.Random
	{
		ran = crypto.New(crypto.Config{})
	}

	var xys plotter.XYs
	for i := 0; i < 10000; i++ {
		var rat float32
		{
			rat = ran.F32()
		}

		var prb float32
		{
			prb = ran.F32()
		}

		if act.Act(rat, prb) {
			xys = append(xys, plotter.XY{X: float64(rat), Y: float64(prb)})
		}
	}

	{
		writer("harrel.svg", buffer(xys))
	}
}

func buffer(xys plotter.XYs) []byte {
	var err error

	var plo *plot.Plot
	{
		plo = plot.New()
	}

	{
		plo.X.Label.Text = "rat"
		plo.X.Min = 0
		plo.X.Max = 1
		plo.Y.Label.Text = "prb"
		plo.Y.Min = 0
		plo.Y.Max = 1
	}

	var sca *plotter.Scatter
	{
		sca, err = plotter.NewScatter(xys)
		if err != nil {
			panic(err)
		}
	}

	{
		sca.GlyphStyle.Shape = draw.CrossGlyph{}
	}

	{
		plo.Add(sca)
	}

	var byt []byte
	{
		wri, err := plo.WriterTo(5*vg.Inch, 5*vg.Inch, "svg")
		if err != nil {
			panic(err)
		}

		buf := bytes.NewBuffer([]byte{})

		_, err = wri.WriteTo(buf)
		if err != nil {
			panic(err)
		}

		byt = buf.Bytes()
	}

	return byt
}

func writer(pat string, byt []byte) {
	err := os.WriteFile(pat, byt, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Rendered scatter plot %s\n", pat)
}
```
