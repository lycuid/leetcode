// https://leetcode.com/problems/unique-paths/
package main

func uniquePaths(rows int, cols int) int {
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < cols; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[rows-1][cols-1]
}

func main() {}
