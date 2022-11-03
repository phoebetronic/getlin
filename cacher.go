package getlin

type Cacher interface {
	Add(Binary)
	Del()
	Lat() Vector
	Log()
	Out(int) []Binary
	Upd(Binary)
	Vec(int) Vector
}
