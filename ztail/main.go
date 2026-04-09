package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		os.Exit(1)
	}

	if args[0] != "-c" {
		os.Exit(1)
	}

	numBytes := 0
	for _, char := range args[1] {
		if char < '0' || char > '9' {
			os.Exit(1)
		}
		numBytes = numBytes*10 + int(char-'0')
	}

	files := args[2:]
	hasError := false
	multipleFiles := len(files) > 1
	firstOutput := true

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			hasError = true
			continue
		}

		if firstOutput && hasError {
			fmt.Println()
		}

		if multipleFiles {
			if !firstOutput {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", filename)
		}

		firstOutput = false

		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			file.Close()
			hasError = true
			continue
		}

		fileSize := fileInfo.Size()
		var offset int64
		if int64(numBytes) >= fileSize {
			offset = 0
		} else {
			offset = fileSize - int64(numBytes)
		}

		_, err = file.Seek(offset, 0)
		if err != nil {
			fmt.Println(err)
			file.Close()
			hasError = true
			continue
		}

		buffer := make([]byte, numBytes)
		n, _ := file.Read(buffer)

		fmt.Print(string(buffer[:n]))

		file.Close()
	}

	if hasError {
		os.Exit(1)
	}
}
