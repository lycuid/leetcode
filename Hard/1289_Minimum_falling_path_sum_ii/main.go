// https://leetcode.com/problems/minimum-falling-path-sum-ii/
package main

import "math"

func minFallingPathSum(grid [][]int) int {
	var (
		n     = len(grid)
		cache = make([][]int, n+1)
	)
	for i := range cache {
		cache[i] = make([]int, n+1)
		cache[i][0] = math.MaxInt
	}
	grid = append([][]int{make([]int, n)}, grid...)
	for i, post := 1, make([]int, n+1); i <= n; i++ {
		pre := math.MaxInt
		for j := range grid[i] {
			min := pre
			if post[j+1] < min {
				min = post[j+1]
			}
			grid[i][j] += min
			if pre > grid[i-1][j] {
				pre = grid[i-1][j]
			}
			if cache[i][j+1] = grid[i][j]; cache[i][j] < cache[i][j+1] {
				cache[i][j+1] = cache[i][j]
			}
		}
		post[n] = math.MaxInt
		for j := n - 1; j >= 0; j-- {
			if post[j] = grid[i][j]; post[j+1] < post[j] {
				post[j] = post[j+1]
			}
		}
	}
	return cache[n][n]
}

func main() {}
