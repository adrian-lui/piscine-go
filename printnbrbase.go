package piscine

import (
	"github.com/01-edu/z01"
)

var hugeNum bool = false

func PrintNbrBase(nbr int, base string) {
	hugeNum = false
	if !validity(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			nbr++
			hugeNum = true
		}
		nbr *= -1
	}
	var baseLen int = len(base)
	var div int = nbr
	var runeList []rune
	runeList = findDiv(div, baseLen, runeList, base)
	for index := range runeList {
		if hugeNum && (index == len(runeList)-1) {
			z01.PrintRune(runeList[len(runeList)+1-index])
		} else {
			z01.PrintRune(runeList[len(runeList)-1-index])
		}
	}
}

func findDiv(div int, baseLen int, runeList []rune, base string) []rune {
	if div/baseLen != 0 {
		runeList = append(runeList, rune(base[div%baseLen]))
		return findDiv(div/baseLen, baseLen, runeList, base)
	} else {
		runeList = append(runeList, rune(base[div%baseLen]))
		return runeList
	}
}

func validity(base string) bool {
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
	}
	return true
}
