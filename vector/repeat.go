package vector

func Repeat(bit []bool, cou int) []bool {
	var rep []bool

	for i := 0; i < cou; i++ {
		rep = append(rep, bit...)
	}

	return rep
}
