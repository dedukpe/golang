package main

import "os"

func main() {
	name := os.Args[0]

	// find last slash '/' or backslash '\'
	last := -1
	for i := 0; i < len(name); i++ {
		if name[i] == '/' || name[i] == '\\' {
			last = i
		}
	}

	if last != -1 {
		name = name[last+1:]
	}

	// add newline
	name += "\n"

	_, err := os.Stdout.Write([]byte(name))
	if err != nil {
		os.Exit(1) // exit with error if write fails
	}
}
