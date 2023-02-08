package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := os.Args[1]
	file, _ := os.Open(fileName)
	content := make([]byte, 14)
	file.Read(content)
	fmt.Println(string(content))
	file.Close()
}
