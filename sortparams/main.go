package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var list []string
	for _, arg := range os.Args[1:] {
		if len(list) == 0 {
			list = append(list, arg)
		} else {
			for index, word := range list {
				if arg > word {
					list = append(list[:index+1], list[index:]...)
					list[index] = arg
					break
				}
			}
		}
	}
	for i := len(list) - 1; i >= 0; i-- {
		for _, letter := range list[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
