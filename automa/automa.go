package automa

type Automa struct {
	poi float32
}

func New(con Config) *Automa {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	return &Automa{
		poi: con.Poi,
	}
}

func (a *Automa) Add(del float32) {
	{
		a.poi += del
	}

	if a.poi >= 1 {
		a.poi = 1
	}
}

func (a *Automa) Exc() bool {
	return a.poi <= 0.5
}

func (a *Automa) Hig() bool {
	return a.poi >= 1
}

func (a *Automa) Inc() bool {
	return a.poi > 0.5
}

func (a *Automa) Low() bool {
	return a.poi <= 0
}

func (a *Automa) Poi() float32 {
	return a.poi
}

func (a *Automa) Red(del float32) {
	{
		a.poi -= del
	}

	if a.poi <= 0 {
		a.poi = 0
	}
}
