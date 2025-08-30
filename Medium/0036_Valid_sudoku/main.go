// https://leetcode.com/problems/valid-sudoku/
package main

func isValidSudoku(board [][]byte) bool {
	for i := range board {
		for j := range board[i] {
			if ch := board[i][j]; ch != '.' {
				board[i][j] = '.'
				for k := 0; k < 9; k++ {
					if board[i][k] == ch {
						return false
					}
					if board[k][j] == ch {
						return false
					}
					if board[i/3*3+(k/3)][j/3*3+(k%3)] == ch {
						return false
					}
				}
				board[i][j] = ch
			}
		}
	}
	return true
}

func main() {}
