// https://leetcode.com/problems/word-search/
package main

func Aux(y, x int, word string, grid [][]byte) bool {
	if len(word) == 1 {
		return true
	}
	grid[y][x] = '#'
	for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		dy, dx := y+d[0], x+d[1]
		if dy >= 0 && dy < len(grid) && dx >= 0 && dx < len(grid[0]) {
			if grid[dy][dx] == word[1] && Aux(dy, dx, word[1:], grid) {
				return true
			}
		}
	}
	grid[y][x] = word[0]
	return false
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] && Aux(i, j, word, board) {
				return true
			}
		}
	}
	return false
}

func main() {}
