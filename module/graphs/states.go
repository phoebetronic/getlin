package graphs

func (m *Module) States() []float32 {
	var sta []float32

	for _, x := range m.mpr.All() {
		for _, y := range x {
			sta = append(sta, y.States()...)
		}
	}

	return sta
}
