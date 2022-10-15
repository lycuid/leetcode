// https://leetcode.com/problems/string-compression-ii/
package main

func getLengthOfOptimalCompression(s string, k int) int {
	var cache [110][110]int
	n := len(s)
	for i := range cache {
		for j := range cache[i] {
			cache[i][j] = 1e5
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
					d := 0
					if count >= 100 {
						d = 3
					} else if count >= 10 {
						d = 2
					} else if count >= 2 {
						d = 1
					}
					if m := cache[l-1][j-del] + 1 + d; m < cache[i][j] {
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
