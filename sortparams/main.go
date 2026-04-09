package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[j] < args[i] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
