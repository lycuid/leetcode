// https://leetcode.com/problems/longest-common-subsequence/
package main

func longestCommonSubsequence(text1 string, text2 string) int {
	cache, m, n := make([][]int, len(text1)+1), len(text1), len(text2)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if cache[i][j] = cache[i][j-1]; cache[i][j] < cache[i-1][j] {
				cache[i][j] = cache[i-1][j]
			}
			if text1[i-1] == text2[j-1] {
				cache[i][j] = cache[i-1][j-1] + 1
			}
		}
	}
	return cache[m][n]
}

func main() {}
