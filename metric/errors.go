package metric

type errors struct {
	cou float32
	los float32
}

func (e *errors) Mae() float32 {
	return e.los / e.cou
}
