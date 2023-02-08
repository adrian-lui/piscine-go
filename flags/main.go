package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Print("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func sort(strPrint string) {
	var list []string
	for _, arg := range strPrint {
		if len(list) == 0 {
			list = append(list, string(' '))
		}
		for index, word := range list {
			if string(arg) >= word {
				list = append(list[:index+1], list[index:]...)
				list[index] = string(arg)
				break
			}
		}
	}
	for i := len(list) - 2; i >= 0; i-- {
		for _, letter := range list[i] {
			z01.PrintRune(letter)
		}
	}
}

func main() {
	var newStr string
	switch len(os.Args) {
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printHelp()
		} else {
			fmt.Print(os.Args[1])
		}
	case 3:
		if os.Args[1] == "--order" || os.Args[1] == "-o" {
			sort(os.Args[2])
		} else {
			fmt.Print(os.Args[2])
			for index, match := range os.Args[1] {
				if match == '=' {
					fmt.Print(os.Args[1][index+1:])
				}
			}
		}
	case 4:
		for index, match := range os.Args[1] {
			if match == '=' {
				newStr = os.Args[3] + os.Args[1][index+1:]
			}
		}
		sort(newStr)
	default:
		printHelp()
	}
	z01.PrintRune('\n')
}
