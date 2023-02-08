package piscine

func Abort(a, b, c, d, e int) int {
	intList := []int{a, b, c, d, e}

	for i := 0; i < len(intList)-1; i++ {
		for j := 1 + i; j < len(intList); j++ {
			if intList[i] > intList[j] {
				intList[i], intList[j] = intList[j], intList[i]
			}
		}
	}
	return intList[len(intList)/2]
}
