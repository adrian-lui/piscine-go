package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var upper bool = false
	var validity string
	var start int
	if len(os.Args) == 1 {
		return
	}
	if os.Args[1] == "--upper" {
		upper = true
		start = 2
	} else {
		start = 1
	}
	for i := start; i < len(os.Args); i++ {
		word := os.Args[i]
		validity = checkValid(word)
		if validity == "invalid" {
			z01.PrintRune(' ')
			continue
		}
		if validity == "oneDigit" {
			if upper {
				z01.PrintRune(rune(int(word[0]) + 17 - 1))
			} else {
				z01.PrintRune(rune(int(word[0]) + 17 - 1 + 32))
			}
			continue
		}
		if validity == "twoDigits" {
			var sum int
			if upper {
				sum = (int(word[0])-48)*10 + 17 - 1 + (int(word[1]))
			} else {
				sum = (int(word[0])-48)*10 + 17 - 1 + (int(word[1]) + 32)
			}
			z01.PrintRune(rune(sum))
			continue
		}
	}
	z01.PrintRune('\n')
}

func checkValid(argument string) string {
	if len(argument) > 2 {
		return "invalid"
	}
	if len(argument) == 1 {
		if argument[0] < '0' || argument[0] > '9' {
			return "invalid"
		} else {
			return "oneDigit"
		}
	}
	if len(argument) == 2 {
		if (argument[0] != '1' && argument[0] != '2') || (argument[1] < '0' && argument[1] > '9') {
			return "invalid"
		} else if argument[1] > '6' && argument[0] == '2' {
			return "invalid"
		} else {
			return "twoDigits"
		}
	}
	return "nil"
}
