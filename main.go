package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	if fileName == "" {
		fmt.Println("Please provide a file")
	} else {
		fmt.Println(fileName)
	}
}
