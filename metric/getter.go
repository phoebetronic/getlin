package metric

import "github.com/phoebetron/getlin"

type getter struct {
	sha *shared
}

func (g *getter) Err() getlin.Errors {
	return g.sha.err
}

func (g *getter) Sta() getlin.States {
	return g.sha.sta
}
