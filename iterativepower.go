package piscine

func IterativePower(nb int, power int) int {
	sum := 1
	if power < 0 {
		return 0
	} else {
		for i := power; i >= 1; i-- {
			sum *= nb
		}
		return sum
	}
}
