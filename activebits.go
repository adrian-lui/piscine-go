package piscine

func ActiveBits(n int) int {
	var count int = 0
	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}
