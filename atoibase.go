package piscine

func valid9(base string) bool {
	for index1, match1 := range base {
		if match1 == '+' || match1 == '-' {
			return false
		}
		for index2, match2 := range base {
			if match1 == match2 && index1 != index2 {
				return false
			}
		}
	}
	if len(base) < 2 {
		return false
	} else {
		return true
	}
}

func RecursivePower2(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	} else {
		return nb * RecursivePower2(nb, power-1)
	}
}

func AtoiBase(s string, base string) int {
	if !valid9(base) {
		return 0
	}
	var sum int = 0
	for i := 0; i <= len(s)-1; i++ {
		for index, match := range base {
			if match == rune(s[i]) {
				sum += RecursivePower2(len(base), len(s)-i-1) * index
			}
		}
	}
	return sum
}
