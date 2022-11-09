package stat32

func Absf32(f float32) float32 {
	if f < 0 {
		return -f
	}

	return f
}
