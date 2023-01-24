// https://leetcode.com/problems/snakes-and-ladders/
package main

func snakesAndLadders(board [][]int) int {
	n := len(board)
	board_value := func(i int) int {
		y, x := i/n, i%n
		if y%2 == 1 {
			x = n - 1 - x
		}
		return board[n-1-y][x]
	}
	weight, cache := make([]int, n*n), []int{0}
	for i := 1; i < n*n; i++ {
		weight[i] = -1
	}
	for ; len(cache) > 0; cache = cache[1:] {
		root := cache[0]
		for i := root + 1; i <= root+6; i++ {
			if i >= n*n {
				break
			}
			child := i
			if index := board_value(child); index != -1 {
				child = index - 1
			}
			if weight[child] == -1 {
				weight[child] = weight[root] + 1
				cache = append(cache, child)
			}
		}
	}
	return weight[n*n-1]
}

func main() {}
