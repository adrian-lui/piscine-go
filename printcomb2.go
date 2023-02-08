package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for m := '0'; m <= '9'; m++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					if i >= k {
						k = i
						if j >= m {
							m = j + 1
						}
					}
					if m > '9' {
						m = '0'
						k++
					}
					z01.PrintRune(k)
					z01.PrintRune(m)
					if i == '9' && j == '8' {
						z01.PrintRune('\n')
						return
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
