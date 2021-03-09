package bomberman

import (
	"strings"
)

// Complete the bomberMan function below.
func bomberMan(n int32, grid []string) []string {
	if n < 2 {
		return grid
	}
	if n%2 == 0 {
		for i := range grid {
			grid[i] = strings.Repeat("O", len(grid[i]))
		}
		return grid
	}

	board := make([][]byte, len(grid))
	for i := range grid {
		board[i] = []byte(grid[i])
	}
	rowlen := len(grid[0])
	m := len(grid)

	var sol [2][]string
	for round := 3; round <= 5; round += 2 {
		// Pass 1. mark neighbouring nodes as 'N'
		for i := range board {
			for j := range board[i] {
				if board[i][j] == 'O' {
					for _, pos := range [][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
						if pos[0] < 0 || pos[1] < 0 || pos[0] >= m || pos[1] >= rowlen || board[pos[0]][pos[1]] == 'O' {
							continue
						}
						board[pos[0]][pos[1]] = 'N'
					}
				}
			}
		}

		// Pass 2. invert all 'N' and 'O' => '.' and '.' => 'O'
		for i := range board {
			for j := range board[i] {
				if board[i][j] == 'O' || board[i][j] == 'N' {
					board[i][j] = '.'
				} else {
					board[i][j] = 'O'
				}
			}
		}

		sol[(round-2)/2] = make([]string, len(grid))
		for i := range board {
			sol[(round-2)/2][i] = string(board[i])
		}
	}

	if (n/2)%2 == 0 {
		return sol[1]
	}

	return sol[0]
}
