package vector

func ToBool(lis ...int) []bool {
	bol := []bool{}

	for _, x := range lis {
		if x == 0 {
			bol = append(bol, false)
		}
		if x == 1 {
			bol = append(bol, true)
		}
	}

	return bol
}
