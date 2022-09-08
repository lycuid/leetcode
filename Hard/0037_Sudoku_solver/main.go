// https://leetcode.com/problems/sudoku-solver/
package main

func ValidBoard(x, y int, ch byte, board *[][]byte) bool {
	for i := 0; i < 9; i++ {
		hx, hy := x, i
		if (*board)[hx][hy] == ch {
			return false
		}
		vx, vy := i, y
		if (*board)[vx][vy] == ch {
			return false
		}
		bx, by := ((x/3)*3)+(i/3), ((y/3)*3)+(i%3)
		if (*board)[bx][by] == ch {
			return false
		}
	}
	return true
}

func Solve(i, j int, board *[][]byte) bool {
	if i == 9 {
		return true
	}
	if j == 9 {
		return Solve(i+1, 0, board)
	}
	if (*board)[i][j] != '.' {
		return Solve(i, j+1, board)
	}
	for ch := byte('1'); ch <= '9'; ch++ {
		if ValidBoard(i, j, ch, board) {
			(*board)[i][j] = ch
			if Solve(i, j+1, board) {
				return true
			}
		}
		(*board)[i][j] = '.'
	}
	return false
}

func solveSudoku(board [][]byte) {
	Solve(0, 0, &board)
}

func main() {}
