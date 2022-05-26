package main

import (
	"fmt"
	"io"
	"os"
)

// in order to successfully run this program, do "go run main.go <file-name>"
func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error at reading file", fileName)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
