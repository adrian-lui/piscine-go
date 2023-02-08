package piscine

func BasicAtoi(s string) int {
	var sum int = 0
	for index, x := range s {
		digit := int(x) - 48
		power := (len(s) - index - 1)
		multiplier := 1
		for i := power; i > 0; i-- {
			multiplier *= 10
		}
		sum += digit * multiplier
	}
	return sum
}
