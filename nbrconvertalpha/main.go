package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	if len(args) == 0 {
		// no arguments, print nothing, no newline
		return
	}

	for _, arg := range args {
		n := atoi(arg)
		if n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		var letter rune = rune('a' + n - 1)
		if upper {
			letter = rune('A' + n - 1)
		}
		z01.PrintRune(letter)
	}

	// Print newline only if we printed something
	z01.PrintRune('\n')
}

func atoi(s string) int {
	if len(s) == 0 {
		return -1
	}

	num := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return -1
		}
		num = num*10 + int(r-'0')
	}
	return num
}
