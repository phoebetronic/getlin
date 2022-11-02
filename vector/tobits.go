package vector

func ToBits(raw ...bool) []int {
	var bit []int

	for _, x := range raw {
		if x {
			bit = append(bit, 1)
		} else {
			bit = append(bit, 0)
		}
	}

	return bit
}
