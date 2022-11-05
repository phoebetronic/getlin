# render

Module architectures of `graphs` Modules can be rendered using [Graphviz] and
the [DOT language]. Reading from left to right, in our example we design a
learnable system that requires an input Vector of 512 bits, and produces an
output Vector of 2 bits. The vertically alligned layers are concatenated and
connected with each Module of the following layer. Modules are described as
records, stating their respective Module names at the top. Graphs of Modules
form arbitrarily nested clusters. Below the left-most cluster produces a
concatenated output Vector of 4 times 16 bits. Each and every of the following
Modules receives this concatenated output Vector as individual input vector for
their respective clusters, forwarding bits and their transformations deeper and
deeper into the system. The following figure can be rendered by running the code
below via `go run main.go`.

[Graphviz]: https://graphviz.org
[DOT language]: https://graphviz.org/doc/info/lang.html

### figure

![Render](/assets/render.svg)

### code

```go
package main

import (
	"fmt"
	"os"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/mapper"
	"github.com/phoebetron/getlin/module/graphs"
	"github.com/phoebetron/getlin/module/linear"
	"github.com/phoebetron/getlin/module/voting"
	"github.com/phoebetron/getlin/render"
)

func main() {
	var mod getlin.Module
	{
		mod = graphs.New(graphs.Config{
			Mpr: mapper.New(mapper.Config{
				Mod: [][]getlin.Module{
					{
						linear.New(linear.Config{Inp: 512, Out: 32, Sta: 32}),
						linear.New(linear.Config{Inp: 512, Out: 32, Sta: 32}),
					},
					{
						linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
						linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
						linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
						linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
					},
					{
						graphs.New(graphs.Config{Mpr: mapper.New(mapper.Config{Mod: [][]getlin.Module{
							{
								linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
								linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
							},
							{voting.New(voting.Config{Inp: 32})},
						}})}),
						graphs.New(graphs.Config{Mpr: mapper.New(mapper.Config{Mod: [][]getlin.Module{
							{
								linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
								linear.New(linear.Config{Inp: 64, Out: 16, Sta: 32}),
							},
							{voting.New(voting.Config{Inp: 32})},
						}})}),
					},
				},
			}),
		})
	}

	{
		writer("render.svg", render.Render(render.Create(mod.Mapper())))
	}
}

func writer(pat string, byt []byte) {
	err := os.WriteFile(pat, byt, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Rendered module architecture %s\n", pat)
}
```
