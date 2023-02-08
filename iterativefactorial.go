package piscine

func IterativeFactorial(nb int) int {
	sum := 1
	if nb < 0 || nb > 63 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			sum *= i
		}
	}
	return sum
}
