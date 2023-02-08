package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var sort []int
	for n > 0 {
		sort = append(sort, n%10)
		n = n / 10
	}
	sort = sorting(sort)
	for _, num := range sort {
		z01.PrintRune(rune(num + 48))
	}
}

func sorting(sort []int) []int {
	var sorted []int
	for small := 0; small <= 9; small++ {
		for _, num := range sort {
			if small == num {
				sorted = append(sorted, small)
			}
		}
	}
	return sorted
}
