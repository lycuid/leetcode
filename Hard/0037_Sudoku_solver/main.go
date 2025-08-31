// https://leetcode.com/problems/sudoku-solver/
package main

func isValid(board [][]byte, i, j int, ch byte) bool {
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
	return true
}

func Aux(board [][]byte, index int) bool {
	if index >= 81 {
		return true
	}
	if board[index/9][index%9] != '.' {
		return Aux(board, index+1)
	}
	for i := 1; i <= 9; i++ {
		if ch := byte(i) + '0'; isValid(board, index/9, index%9, ch) {
			if board[index/9][index%9] = ch; Aux(board, index+1) {
				return true
			}
			board[index/9][index%9] = '.'
		}
	}
	return false
}

func solveSudoku(board [][]byte) {
	Aux(board, 0)
}

func main() {}
