package automa

import "math/rand"

type Automa struct {
	poi float32
	sta float32
}

func New(con Config) *Automa {
	{
		con.Verify()
	}

	return &Automa{
		sta: float32(con.Sta),
	}
}

func (a *Automa) Add(cou float32) {
	if a.poi >= +a.sta {
		return
	}

	{
		a.poi += cou
	}

	if a.poi >= +a.sta {
		a.poi = +a.sta
	}
}

func (a *Automa) Exc() bool {
	return a.poi < 0
}

func (a *Automa) Inc() bool {
	return a.poi > 0
}

func (a *Automa) Ini(dir int) {
	if dir == -1 {
		a.poi = -float32(rand.Intn(int(a.sta)) + 1)
	}

	if dir == +1 {
		a.poi = +float32(rand.Intn(int(a.sta)) + 1)
	}
}

func (a *Automa) Neu() bool {
	return a.poi == 0
}

func (a *Automa) Poi() float32 {
	return a.poi
}

func (a *Automa) Rat() float32 {
	return abs(a.poi) / a.sta
}

func (a *Automa) Rem(cou float32) {
	if a.poi <= -a.sta {
		return
	}

	{
		a.poi -= cou
	}

	if a.poi <= -a.sta {
		a.poi = -a.sta
	}
}

func abs(f float32) float32 {
	if f < 0 {
		return -f
	}

	return f
}
