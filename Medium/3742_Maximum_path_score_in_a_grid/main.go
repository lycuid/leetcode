// https://leetcode.com/problems/maximum-path-score-in-a-grid/
package main

func maxPathScore(grid [][]int, k int) int {
	cost := func(i int) int {
		return (i + 1) >> 1
	}
	m, n := len(grid), len(grid[0])
	cache := make([][][]int, m+1)
	for i := range cache {
		cache[i] = make([][]int, n+1)
		for j := range cache[i] {
			cache[i][j] = make([]int, k+1)
			for _k := range cache[i][j] {
				cache[i][j][_k] = -1
			}
		}
	}
	cache[1][1][cost(grid[0][0])] = grid[0][0]
	for i := range grid {
		for j := range grid[i] {
			c := cost(grid[i][j])
			for _i, score := range cache[i][j+1] {
				if _k := c + _i; _k <= k && score != -1 {
					cache[i+1][j+1][_k] = max(cache[i+1][j+1][_k], score+grid[i][j])
				}
			}
			for _i, score := range cache[i+1][j] {
				if _k := c + _i; _k <= k && score != -1 {
					cache[i+1][j+1][_k] = max(cache[i+1][j+1][_k], score+grid[i][j])
				}
			}
		}
	}
	res := -1
	for _, score := range cache[m][n] {
		res = max(res, score)
	}
	return res
}

func main() {}
