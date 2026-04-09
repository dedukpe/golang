package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	// Iterate over all arguments except the first one (program name)
	for i := 1; i < len(args); i++ {
		for _, r := range args[i] {
			z01.PrintRune(r) // print each rune of the argument
		}
		z01.PrintRune('\n') // print newline after each argument
	}
}
