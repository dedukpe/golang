package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digitsCount [10]int

	for n > 0 {
		d := n % 10
		digitsCount[d]++
		n /= 10
	}

	for i := 0; i <= 9; i++ {
		for j := 0; j < digitsCount[i]; j++ {
			z01.PrintRune(rune(i + '0'))
		}
	}
}
