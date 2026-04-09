package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// Join args with space
	var joined []rune
	for i, arg := range args {
		for _, r := range arg {
			joined = append(joined, r)
		}
		if i != len(args)-1 {
			joined = append(joined, ' ')
		}
	}

	// Collect vowels and their positions
	var vowels []rune
	var vowelPositions []int

	for i, r := range joined {
		if isVowel(r) {
			vowels = append(vowels, r)
			vowelPositions = append(vowelPositions, i)
		}
	}

	// If no vowels, print arguments as is with spaces
	if len(vowels) == 0 {
		for i, arg := range args {
			for _, r := range arg {
				z01.PrintRune(r)
			}
			if i != len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// Reverse vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Replace vowels in joined string with reversed vowels
	for i, pos := range vowelPositions {
		joined[pos] = vowels[i]
	}

	// Print the resulting string
	for _, r := range joined {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
