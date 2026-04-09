package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var comb [9]rune
	printCombNRec(n, 0, 0, comb, false)
	z01.PrintRune('\n')
}

func printCombNRec(n, index, start int, comb [9]rune, printed bool) bool {
	if index == n {
		if printed {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for i := 0; i < n; i++ {
			z01.PrintRune(comb[i] + '0')
		}
		return true
	}

	for digit := start; digit <= 9; digit++ {
		comb[index] = rune(digit)
		printed = printCombNRec(n, index+1, digit+1, comb, printed)
	}
	return printed
}
