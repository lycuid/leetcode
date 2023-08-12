// https://leetcode.com/problems/unique-paths-ii/
package main

func uniquePathsWithObstacles(grid [][]int) int {
	m, n := len(grid)-1, len(grid[0])-1
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = -1
			}
		}
	}
	if grid[0][0] != -1 {
		grid[0][0] = 1
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if grid[i][j] != -1 {
				if i > 0 && grid[i-1][j] != -1 {
					grid[i][j] += grid[i-1][j]
				}
				if j > 0 && grid[i][j-1] != -1 {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}
	if grid[m][n] == -1 {
		grid[m][n] = 0
	}
	return grid[m][n]
}

func main() {}
