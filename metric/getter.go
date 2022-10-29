package metric

import "github.com/phoebetron/getlin"

type getter struct {
	sha *shared
}

func (g *getter) Mat() getlin.Matrix {
	return g.sha.mat
}

func (g *getter) Sta() getlin.States {
	return g.sha.sta
}
