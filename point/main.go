package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

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

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := point{}
	setPoint(&points)
	printStr("x = ")
	PrintNbr(points.x)
	printStr(", y = ")
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
