package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // skip program name

	for i := len(args) - 1; i >= 0; i-- {
		printStr(args[i])
		z01.PrintRune('\n')
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
