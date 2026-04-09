package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printError(err error) {
	msg := "ERROR: " + err.Error() + "\n"
	for _, r := range msg {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for {
			r, _, err := reader.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				os.Exit(1)
			}
			z01.PrintRune(r)
		}
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		buf := make([]byte, 4096)
		for {
			n, err := file.Read(buf)
			if n > 0 {
				for _, b := range buf[:n] {
					z01.PrintRune(rune(b))
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				printError(err)
				os.Exit(1)
			}
		}
		file.Close()
	}
}
