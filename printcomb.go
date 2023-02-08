package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := '1'; j <= '8'; j++ {
			for k := '2'; k <= '9'; k++ {
				z01.PrintRune(i)
				if i >= j {
					j = i + 1
				}
				z01.PrintRune(j)
				if j >= k {
					k = j + 1
				}
				z01.PrintRune(k)
				if i < '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
