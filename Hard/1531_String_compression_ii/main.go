// https://leetcode.com/problems/string-compression-ii/
package main

func digit_count(num int) (n int) {
	if num > 1 {
		for ; num > 0; num /= 10 {
			n++
		}
	}
	return n
}

func getLengthOfOptimalCompression(s string, k int) int {
	cache, n := make([][]int, len(s)+1), len(s)
	for i := range cache {
		cache[i] = make([]int, k+1)
		for j := range cache[i] {
			cache[i][j] = 1e9 + 7
		}
	}
	cache[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			var count, del int
			for l := i; l >= 1; l-- {
				if s[l-1] == s[i-1] {
					count++
				} else {
					del++
				}
				if j-del >= 0 {
					if m := cache[l-1][j-del] + 1 + digit_count(count); m < cache[i][j] {
						cache[i][j] = m
					}
				}
			}
			if j > 0 {
				if m := cache[i-1][j-1]; m < cache[i][j] {
					cache[i][j] = m
				}
			}
		}
	}
	return cache[n][k]
}

func main() {}
