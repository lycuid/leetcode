// https://leetcode.com/problems/game-of-life/
package main

func total_live_neighbours(board [][]int, m, n, x, y int) (total int) {
	if x > 0 {
		total += board[x-1][y]
		if y > 0 {
			total += board[x-1][y-1]
		}
		if y < n-1 {
			total += board[x-1][y+1]
		}
	}
	if x < m-1 {
		total += board[x+1][y]
		if y > 0 {
			total += board[x+1][y-1]
		}
		if y < n-1 {
			total += board[x+1][y+1]
		}
	}
	if y > 0 {
		total += board[x][y-1]
	}
	if y < n-1 {
		total += board[x][y+1]
	}
	return total
}

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	neighbours := make([][]int, m)
	for i := 0; i < m; i++ {
		neighbours[i] = make([]int, n)
		for j := 0; j < n; j++ {
			neighbours[i][j] = total_live_neighbours(board, m, n, i, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if neighbours[i][j] < 2 || neighbours[i][j] > 3 {
				board[i][j] = 0
			} else if neighbours[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

func main() {}
