package getlin

type Cacher interface {
	Add(Binary)
	Del()
	Lat() Vector
	Log()
	Upd(Binary)
	Vec(int) Vector
}
