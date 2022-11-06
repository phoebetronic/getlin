package binary

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (l *Loader) Search() map[int][][2]getlin.Vector {
	return map[int][][2]getlin.Vector{
		0: {
			{
				vector.New(vector.Config{Inp: []uint8{0, 0, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 0, 0, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 0, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 1, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 0, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 1, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 1, 1, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 1, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{1}}),
			},
		},
		1: {
			{
				vector.New(vector.Config{Inp: []uint8{0, 0, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 0, 0, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 0, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 0, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 1, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{0, 1, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 0, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 1, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 0, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 1, 1, 0}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
			{
				vector.New(vector.Config{Inp: []uint8{1, 1, 1, 1}, Out: []uint8{0}}),
				vector.New(vector.Config{Inp: []uint8{1, 1, 0, 1}, Out: []uint8{1}}),
			},
		},
	}
}
