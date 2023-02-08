package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var runes []rune
	if n < 0 {
		z01.PrintRune(45)
	}
	if n == 0 {
		z01.PrintRune(48)
	} else {
		for {
			var remain int = n % 10
			if remain < 0 {
				remain = remain * -1
			}
			runes = append(runes, rune(remain))
			n = n / 10
			if n == 0 {
				break
			}
		}
	}
	for i := len(runes) - 1; i >= 0; i-- {
		z01.PrintRune(runes[i] + 48)
	}
}
