// https://leetcode.com/problems/unique-paths/
package main

func uniquePaths(m int, n int) int {
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	cache[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cache[i][j] += (cache[i-1][j] + cache[i][j-1])
		}
	}
	return cache[m][n]
}

func main() {}
