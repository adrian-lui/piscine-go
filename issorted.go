package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asOrder := true
	deOrder := true

	for num := 0; num <= len(a)-3; num++ {
		if f(a[num], a[num+1]) > 0 {
			asOrder = false
		} else if f(a[num], a[num+1]) < 0 {
			deOrder = false
		}
	}

	if !asOrder && !deOrder {
		return false
	} else {
		return true
	}
}
