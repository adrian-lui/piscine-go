package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, arg := range os.Args[1:] {
		for _, word := range arg {
			z01.PrintRune(rune(word))
		}
		z01.PrintRune('\n')
	}
}
