package binary

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (l *Loader) Search() [][2]getlin.Vector {
	return [][2]getlin.Vector{
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 0, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 1, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 1, 0, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 1, 1, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 1, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 0, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 0, 0, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 0, 1, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 0, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 1, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 1, 0, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 1, 1, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 0, Inp: []float32{1, 1, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 0, Inp: []float32{0, 0, 1, 0}, Out: []float32{1}}),
		},

		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 0, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 0, 0, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 0, 1, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 0, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 1, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 1, 0, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 1, 1, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{0, 1, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 0, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 0, 0, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 0, 1, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 0, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 1, 0}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
		{
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 1, 1}, Out: []float32{0}}),
			vector.New(vector.Config{Cla: 1, Inp: []float32{1, 1, 0, 1}, Out: []float32{1}}),
		},
	}
}
