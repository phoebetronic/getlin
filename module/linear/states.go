package linear

func (m *Module) States() []float32 {
	var sta []float32

	for _, x := range m.cla {
		sta = append(sta, x.States()...)
	}

	return sta
}
