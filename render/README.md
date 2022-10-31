# render

Module architectures of `graphs` modules can be rendered using [Graphviz] and
the [DOT language]. The following figure can be rendered by running the code
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
	"github.com/phoebetron/getlin/module/graphs"
	"github.com/phoebetron/getlin/module/linear"
	"github.com/phoebetron/getlin/render"
)

func main() {
	var mod getlin.Module
	{
		mod = graphs.New(graphs.Config{
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
					linear.New(linear.Config{Inp: 64, Out: 64, Sta: 32}),
				},
			},
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
