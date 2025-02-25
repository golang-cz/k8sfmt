package main

import (
	"fmt"
	"os"

	"github.com/golang-cz/k8sfmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: k8sfmt <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	if err := k8sfmt.PrettifyYAML(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
