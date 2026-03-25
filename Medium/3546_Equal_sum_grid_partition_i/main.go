// https://leetcode.com/problems/equal-sum-grid-partition-i/
package main

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := 1; j < n; j++ {
			grid[i][j] += grid[i][j-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := range grid[i] {
			grid[i][j] += grid[i-1][j]
		}
	}
	for i := range grid {
		if grid[i][n-1]*2 == grid[m-1][n-1] {
			return true
		}
	}
	for _, num := range grid[m-1] {
		if num*2 == grid[m-1][n-1] {
			return true
		}
	}
	return false
}

func main() {}
