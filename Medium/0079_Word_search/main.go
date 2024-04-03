// https://leetcode.com/problems/word-search/
package main

func Aux(y, x int, word string, board [][]byte) bool {
	if len(word) == 1 {
		return true
	}
	board[y][x] = '#'
	for _, d := range [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		dy, dx := y+d[0], x+d[1]
		if dy >= 0 && dy < len(board) && dx >= 0 && dx < len(board[0]) {
			if word[1] == board[dy][dx] && Aux(dy, dx, word[1:], board) {
				return true
			}
		}
	}
	board[y][x] = word[0]
	return false
}

func exist(board [][]byte, word string) bool {
	for y := range board {
		for x := range board[y] {
			if word[0] == board[y][x] && Aux(y, x, word, board) {
				return true
			}
		}
	}
	return false
}

func main() {}
