package linear

func (m *Module) Automa() int {
	return len(m.cla) * m.cla[0].Automa()
}
