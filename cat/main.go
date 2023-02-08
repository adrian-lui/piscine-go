package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		b, _ := ioutil.ReadAll(os.Stdin)
		for _, letter := range b {
			z01.PrintRune(rune(letter))
		}
	} else {
		for index, args := range os.Args {
			if index == 0 {
				continue
			}

			file, err := os.Open(args)
			if err != nil {
				errorMsg := "ERROR: " + err.Error()
				for _, letter := range errorMsg {
					z01.PrintRune(letter)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				stats, _ := os.Stat(args)
				fileSize := stats.Size()
				content := make([]byte, fileSize)
				file.Read(content)
				for _, contentLetter := range content {
					z01.PrintRune(rune(contentLetter))
				}
				file.Close()
			}
		}
	}
}
