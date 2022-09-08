// https://leetcode.com/problems/interleaving-string/
package main

func isInterleave(s1, s2, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	cache := make([][]bool, m+1)
	for i := range cache {
		cache[i] = make([]bool, n+1)
	}
	for i := range cache {
		for j := range cache[i] {
			if i == 0 && j == 0 {
				cache[i][j] = true
			} else if i == 0 {
				cache[i][j] = cache[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				cache[i][j] = cache[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				cache[i][j] = (cache[i-1][j] && s1[i-1] == s3[i+j-1]) || (cache[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return cache[m][n]
}

func main() {}
