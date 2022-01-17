package main

import (
	"fmt"
	"os"
	"simcap/srtest/merkle"
)

func main() {
	data := os.Args[1:]
	if len(data) < 1 {
		fmt.Fprintf(os.Stderr, "usage: go run main.go data_string...\n")
		os.Exit(0)
	}

	tree := merkle.BuildTree(data...)
	fmt.Printf("%d elements given to build tree\n", len(data))
	fmt.Printf("Tree of height %d and hash root %s\n", tree.Height(), tree.Root())
}
