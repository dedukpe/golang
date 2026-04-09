package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	printLine := func(s string) {
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}

	printLine("--insert")
	printLine("  -i")
	printLine("\t This flag inserts the string into the string passed as argument.")
	printLine("--order")
	printLine("  -o")
	printLine("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || contains(args, "--help") || contains(args, "-h") {
		printHelp()
		return
	}

	insertStr := ""
	order := false
	mainStr := ""
	positionalGiven := false

	for _, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if len(arg) >= 0 && (len(arg) == 0 || arg[0] != '-') {
			// Accept empty string
			mainStr = arg
			positionalGiven = true
		}
	}

	if !positionalGiven {
		printHelp()
		return
	}

	mainStr += insertStr

	if order {
		mainStr = orderString(mainStr)
	}

	for _, r := range mainStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func orderString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}

	return string(runes)
}
