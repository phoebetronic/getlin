package metric

type getter struct {
	sha *shared
}

func (g *getter) Mat() Matrix {
	return g.sha.mat
}

func (g *getter) Sta() States {
	return g.sha.sta
}
