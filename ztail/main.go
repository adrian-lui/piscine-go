package main

import (
	"fmt"
	"os"
)

func tail(fileNames string, num int) {
	file, _ := os.Open(fileNames)
	stats, _ := os.Stat(fileNames)
	fileSize := stats.Size()
	content := make([]byte, fileSize)
	file.Read(content)
	if num > len(content) {
		fmt.Println(string(content[0 : len(content)-1]))
	} else {
		fmt.Println(string(content[len(content)-num : len(content)-1]))
	}
	file.Close()
}

func Atoi(s string) int {
	var sum int = 0
	var signed bool = false
	var neg bool = false
	for index, x := range s {
		if x == '+' || x == '-' {
			if signed {
				return 0
			}
			if x == '-' {
				neg = true
			}
			signed = true
			continue
		}
		if x != '+' || x != '-' {
			if x > '9' || x < '0' {
				return 0
			}
		}
		digit := int(x) - 48
		power := (len(s) - index - 1)
		multiplier := 1
		for i := power; i > 0; i-- {
			multiplier *= 10
		}
		sum += digit * multiplier
		signed = true
	}
	if neg {
		sum *= -1
	}
	return sum
}

func main() {
	tailNum := Atoi(os.Args[2])
	var hasError bool = false
	for index, fileNames := range os.Args {
		if index <= 2 {
			continue
		}
		_, err := os.Open(fileNames)
		if err != nil {
			fmt.Println(err.Error())
			hasError = true
			continue
		}
		if index != 3 {
			fmt.Println()
		}
		if len(os.Args) != 4 {
			fmt.Printf("==> %v <==\n", fileNames)
		}
		tail(os.Args[index], tailNum)
	}
	if hasError {
		os.Exit(1)
	}
}
