package main

import "github.com/01-edu/z01"

func main() {
	output := []rune{
		120, 32, 61, 32, 52, 50, 44, 32, 121, 32, 61, 32, 50, 49, 10,
	}

	for _, v := range output {
		z01.PrintRune(v)
	}
}
