package getlin

type Mapper interface {
	Abo(Module) []Module
	All() [][]Module
	Bel(Module) []Module
	Ind(Module) int
}
