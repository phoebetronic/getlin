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

	if a.poi >= 2*a.sta {
		a.poi = (2 * a.sta)
	}
}

func (a *Automa) Exc() bool {
	return a.poi <= a.sta
}

func (a *Automa) Hig() bool {
	return a.poi >= (2 * a.sta)
}

func (a *Automa) Inc() bool {
	return a.poi > a.sta
}

func (a *Automa) Low() bool {
	return a.poi <= 1
}

func (a *Automa) Poi() int {
	return a.poi
}

func (a *Automa) Red() {
	{
		a.poi -= 1
	}

	if a.poi <= 1 {
		a.poi = 1
	}
}
