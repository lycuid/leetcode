// https://leetcode.com/problems/longest-palindromic-subsequence/
package main

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestPalindromeSubseq(s string) int {
	cache, n := [2][]int{make([]int, len(s)+1), make([]int, len(s)+1)}, len(s)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == s[n-j] {
				cache[1][j] = 1 + cache[0][j-1]
			} else {
				cache[1][j] = Max(cache[1][j-1], cache[0][j])
			}
		}
		cache[0], cache[1] = cache[1], cache[0]
	}
	return cache[0][n]
}

func main() {}
