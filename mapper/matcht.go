package mapper

func matcht(lay []int, out int, ind int) [2]int {
	if out == 1 {
		return [2]int{0, 1}
	}

	var lef int
	var mid int
	var sum int
	for i, x := range lay {
		if i < ind {
			lef += x
		}

		if i == ind {
			mid = x
		}

		{
			sum += x
		}
	}

	if sum == out {
		return [2]int{lef, lef + mid}
	}

	if len(lay)%2 != sum%2 {
		return [2]int{}
	}

	if out > sum {
		return [2]int{}
	}

	var sta int
	{
		sta = int(float32(lef) / float32(sum) * float32(out))
	}

	var end int
	{
		end = int(float32(mid)/float32(sum)*float32(out) + 0.5)
	}

	return [2]int{sta, sta + end}
}
