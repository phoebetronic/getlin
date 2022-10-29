package getlin

type Cacher interface {
	Create(Vector)
	Delete()
	Latest() Linker
	Update(Vector)
}

type Linker struct {
	Inp Vector
	Out Vector
}
