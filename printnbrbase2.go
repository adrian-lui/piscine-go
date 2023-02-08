package piscine

import (
	"github.com/01-edu/z01"
)

var hugeNum2 bool = false

func PrintNbrBase2(nbr int, base string) string {
	hugeNum = false
	if !validity2(base) {
		return "NV"
	}
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			nbr++
			hugeNum2 = true
		}
		nbr *= -1
	}
	var baseLen int = len(base)
	var div int = nbr
	var runeList []rune
	runeList = findDiv2(div, baseLen, runeList, base)
	var toPrint string
	for index := range runeList {
		if hugeNum2 && (index == len(runeList)-1) {
			toPrint += string(runeList[len(runeList)+1-index])
		} else {
			toPrint += string(runeList[len(runeList)-1-index])
		}
	}
	return toPrint
}

func findDiv2(div int, baseLen int, runeList []rune, base string) []rune {
	if div/baseLen != 0 {
		runeList = append(runeList, rune(base[div%baseLen]))
		return findDiv2(div/baseLen, baseLen, runeList, base)
	} else {
		runeList = append(runeList, rune(base[div%baseLen]))
		return runeList
	}
}

func validity2(base string) bool {
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
	return len(base) >= 2
}
