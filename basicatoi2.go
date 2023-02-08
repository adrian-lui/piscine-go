package piscine

func BasicAtoi2(s string) int {
	var sum int = 0
	for index, x := range s {
		if x > '9' || x < '0' {
			return 0
		}
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
