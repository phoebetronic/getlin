package stat32

func Minf32(a float32, b float32) float32 {
	if a < b {
		return a
	}

	return b
}
