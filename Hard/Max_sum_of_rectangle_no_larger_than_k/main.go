// https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/
package main

func maxSumSubmatrix(matrix [][]int, k int) int {
	r, c := len(matrix), len(matrix[0])
	dp := make([][]int, r+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			dp[i][j] = matrix[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	var ret int = -1e5
	for i1 := 1; i1 <= r; i1++ {
		for i2 := i1; i2 <= r; i2++ {
			for j1 := 1; j1 <= c; j1++ {
				for j2 := j1; j2 <= c; j2++ {
					if s := dp[i2][j2] - dp[i1-1][j2] - dp[i2][j1-1] + dp[i1-1][j1-1]; s > ret && s <= k {
						ret = s
					}
				}
			}
		}
	}
	return ret
}

func main() {}
