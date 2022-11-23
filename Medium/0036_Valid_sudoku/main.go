// https://leetcode.com/problems/valid-sudoku/
package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var horizontal, vertical, box [9]bool
		for j := 0; j < 9; j++ {
			if hy, hx := i, j; board[hy][hx] != '.' {
				if horizontal[board[hy][hx]-'1'] {
					return false
				}
				horizontal[board[hy][hx]-'1'] = true
			}
			if vy, vx := j, i; board[vy][vx] != '.' {
				if vertical[board[vy][vx]-'1'] {
					return false
				}
				vertical[board[vy][vx]-'1'] = true
			}
			if by, bx := ((i%3)*3)+(j/3), ((i/3)*3)+(j%3); board[by][bx] != '.' {
				if box[board[by][bx]-'1'] {
					return false
				}
				box[board[by][bx]-'1'] = true
			}
		}
	}
	return true
}

func main() {}
