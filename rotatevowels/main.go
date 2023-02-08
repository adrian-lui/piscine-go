package main

import (
	"os"

	"github.com/01-edu/z01"
)

func findVowels(newStr string) []int {
	var vowelIndex []int
	for index, match := range newStr {
		if match == 'a' || match == 'e' || match == 'i' || match == 'o' || match == 'u' || match == 'A' || match == 'E' || match == 'I' || match == 'O' || match == 'U' {
			vowelIndex = append(vowelIndex, index)
		}
	}
	return vowelIndex
}

func main() {
	if len(os.Args) <= 1 {
		z01.PrintRune('\n')
		return
	}
	var newStr string
	for i := 1; i < len(os.Args); i++ {
		newStr += os.Args[i]
		newStr += string(' ')
	}
	vowelIndex := findVowels(newStr)
	var finalStr []rune
	for _, word := range newStr {
		finalStr = append(finalStr, word)
	}
	for k := 0; k < len(vowelIndex)/2; k++ {
		temp := finalStr[vowelIndex[k]]
		finalStr[vowelIndex[k]] = finalStr[vowelIndex[len(vowelIndex)-k-1]]
		finalStr[vowelIndex[len(vowelIndex)-k-1]] = temp
	}
	for _, done := range finalStr {
		z01.PrintRune(done)
	}
	z01.PrintRune('\n')
}
