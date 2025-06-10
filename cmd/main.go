package main

import (
	"fmt"
	"os"

	printing "ascii-art/art"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("usage: go run ./cmd <string>")
		return
	}

	input := args[0]

	printing.Processing(input)
}
