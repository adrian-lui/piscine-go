package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	var count int = 0
	for start != 1 {
		if start%2 == 1 {
			start = start*3 + 1
		} else {
			start /= 2
		}
		count++
	}
	return count
}
