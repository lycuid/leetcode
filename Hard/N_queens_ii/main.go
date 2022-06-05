// https://leetcode.com/problems/n-queens-ii/
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

func Solve(row int, board *[][]bool) (ret int) {
	if row == len(*board) {
		return 1
	}
	for i := 0; i < len(*board); i++ {
		if IsBoardValid(row, i, board) {
			(*board)[row][i] = true
			ret += Solve(row+1, board)
			(*board)[row][i] = false
		}
	}
	return ret
}

func totalNQueens(n int) int {
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}
	return Solve(0, &board)
}

func main() {}
