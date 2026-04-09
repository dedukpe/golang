package piscine

import "github.com/01-edu/z01"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	chars := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if chars[r] {
			return false
		}
		chars[r] = true
	}
	return true
}

func printNbrBaseRecursive(n uint, base string, baseLen uint) {
	if n >= baseLen {
		printNbrBaseRecursive(n/baseLen, base, baseLen)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := uint(len(base))

	if nbr < 0 {
		z01.PrintRune('-')
		var un uint
		if nbr == -2147483648 {
			un = 2147483648
		} else {
			un = uint(-nbr)
		}
		printNbrBaseRecursive(un, base, baseLen)
	} else {
		printNbrBaseRecursive(uint(nbr), base, baseLen)
	}
}
