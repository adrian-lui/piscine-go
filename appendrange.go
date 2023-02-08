package piscine

func AppendRange(min, max int) []int {
	var intSlice []int
	if min >= max {
		return intSlice
	} else {
		for i := min; i < max; i++ {
			intSlice = append(intSlice, i)
		}
	}
	return intSlice
}
