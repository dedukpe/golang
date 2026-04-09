package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(&board, 0)
}

func solve(board *[8]int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}
	for row := 1; row <= 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board *[8]int, col, row int) bool {
	for c := 0; c < col; c++ {
		r := board[c]
		if r == row || abs(r-row) == abs(c-col) {
			return false
		}
	}
	return true
}

func printSolution(board *[8]int) {
	for _, r := range board {
		z01.PrintRune(rune('0' + r))
	}
	z01.PrintRune('\n')
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
