// https://leetcode.com/problems/n-queens/
package main

func IsBoardValid(x, y int, board *[][]bool) bool {
	n := len(*board)
	for i := 0; i < n; i++ {
		if (*board)[x][i] || (*board)[i][y] {
			return false
		}
		if i <= x && i <= y && (*board)[x-i][y-i] {
			return false
		}
		if i <= x && i < n-y && (*board)[x-i][y+i] {
			return false
		}
		if i < n-x && i <= y && (*board)[x+i][y-i] {
			return false
		}
		if i < n-x && i < n-y && (*board)[x+i][y+i] {
			return false
		}
	}
	return true
}

func Translate(board *[][]bool) (ret []string) {
	n := len(*board)
	ret = make([]string, n)
	for i := 0; i < n; i++ {
		b := make([]byte, n)
		for j := 0; j < n; j++ {
			if (*board)[i][j] {
				b[j] = 'Q'
			} else {
				b[j] = '.'
			}
		}
		ret[i] = string(b)
	}
	return ret
}

func Solve(row int, board *[][]bool, ret *[][]string) {
	if row == len(*board) {
		*ret = append(*ret, Translate(board))
		return
	}
	for col := 0; col < len(*board); col++ {
		if IsBoardValid(row, col, board) {
			(*board)[row][col] = true
			Solve(row+1, board, ret)
			(*board)[row][col] = false
		}
	}
}

func solveNQueens(n int) (ret [][]string) {
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}
	Solve(0, &board, &ret)
	return ret
}

func main() {}
