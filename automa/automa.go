package automa

type Automa struct {
	poi int
	sta int
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
		sta: con.Sta,
	}
}

func (a *Automa) Add() {
	{
		a.poi += 1
	}

	if a.poi >= (2*a.sta)+1 {
		a.poi = (2 * a.sta) + 1
	}
}

func (a *Automa) Exc() bool {
	return a.poi <= a.sta
}

func (a *Automa) Hig() bool {
	x := (2 * a.sta) + 1
	return a.poi >= x
}

func (a *Automa) Inc() bool {
	return a.poi > a.sta
}

func (a *Automa) Low() bool {
	return a.poi <= 0
}

func (a *Automa) Poi() int {
	return a.poi
}

func (a *Automa) Red() {
	{
		a.poi -= 1
	}

	if a.poi <= 0 {
		a.poi = 0
	}
}
