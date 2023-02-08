package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	intSlice := make([]int, max-min)
	for i := min; i < max; i++ {
		intSlice[i-min] = i
	}
	return intSlice
}
