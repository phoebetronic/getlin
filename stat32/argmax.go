package stat32

import "math"

func Argmax(lis []float32) int {
	var ind int

	var max float32
	{
		max = -math.MaxFloat32
	}

	for i := range lis {
		if lis[i] > max {
			ind = i
			max = lis[i]
		}
	}

	return ind
}
