package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for index, name := range os.Args[0] {
		if index > 1 {
			z01.PrintRune(rune(name))
		}
	}
	z01.PrintRune('\n')
}
