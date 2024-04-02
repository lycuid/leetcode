// https://leetcode.com/problems/isomorphic-strings/
package main

func isIsomorphic(s string, t string) bool {
	if n := len(s); len(t) == n {
		var cache [2][128]byte
		for i := 0; i < n; i++ {
			if cache[0][s[i]] == 0 && cache[1][t[i]] == 0 {
				cache[0][s[i]], cache[1][t[i]] = t[i], s[i]
			}
			if cache[0][s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}

func main() {}
