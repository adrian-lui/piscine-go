package piscine

func Atoi(s string) int {
	var sum int = 0
	var signed bool = false
	var neg bool = false
	for index, x := range s {
		if x == '+' || x == '-' {
			if signed {
				return 0
			}
			if x == '-' {
				neg = true
			}
			signed = true
			continue
		}
		if x != '+' || x != '-' {
			if x > '9' || x < '0' {
				return 0
			}
		}
		digit := int(x) - 48
		power := (len(s) - index - 1)
		multiplier := 1
		for i := power; i > 0; i-- {
			multiplier *= 10
		}
		sum += digit * multiplier
		signed = true
	}
	if neg {
		sum *= -1
	}
	return sum
}
