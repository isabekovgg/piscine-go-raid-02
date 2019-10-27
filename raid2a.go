package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	sudoku := os.Args[1:]
	//sudoku := []string{"8........", "..36.....", ".7..9.2..", ".5...7...", "....457..", "...1...3.", "..1....68", "..85...1.", ".9....4.."}
	dots := 0
	for _, line := range sudoku {
		for _, num := range line {
			if num == '.' {
				dots++
			}

		}
	}
	Solve(sudoku, 0, 0, dots, 0)
}

func Solve(lines []string, row int, column int, dots int, current int) {
	if current == dots {
		for line := range lines {
			for num := range lines[line] {
				z01.PrintRune(rune(lines[line][num]))

			}
			z01.PrintRune('\n')
		}
	} else {
		line := []byte(lines[row])
		if line[column] == '.' {
			for i := '1'; i <= '9'; i++ {
				line[column] = byte(i)
				lines[row] = string(line)

				if ValidNumber(row, column, lines) {

					if column == 8 {
						Solve(lines, row+1, 0, dots, current+1)
					} else {
						Solve(lines, row, column+1, dots, current+1)
					}
				}
			}
			line[column] = '.'
			lines[row] = string(line)
		} else if column == 8 {
			Solve(lines, row+1, 0, dots, current)
		} else {
			Solve(lines, row, column+1, dots, current)
		}

	}
}

func ValidNumber(row int, column int, lines []string) bool {
	oneinline := true
	oneincolumn := true
	oneinsquare := true
	// one in line
	for index := range lines[row] {
		if oneinline && index != column && lines[row][index] == lines[row][column] {
			oneinline = false
		}
	}
	// one in column
	if oneinline {
		for index, line := range lines {
			if oneincolumn && index != row && lines[row][column] == line[column] {
				oneincolumn = false
			}
		}
		// one in square
		if oneincolumn {
			for i := (row / 3) * 3; i < (row/3)*3+3; i++ {
				for j := (column / 3) * 3; j < (column/3)*3+3; j++ {
					if oneinsquare && i != row || j != column {
						if lines[i][j] == lines[row][column] {
							oneinsquare = false
						}
					}
				}
			}
		}
	}

	if oneinline && oneincolumn && oneinsquare {
		return true
	}
	return false

}
