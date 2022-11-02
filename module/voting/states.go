package voting

func (m *Module) States() []float32 {
	if m.wei != nil {
		return m.wei.States()
	}

	return nil
}
