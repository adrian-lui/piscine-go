package main

import (
	"fmt"
	"os"
)

func main() {
	for _, match := range os.Args {
		if match == "01" || match == "galaxy" || match == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
