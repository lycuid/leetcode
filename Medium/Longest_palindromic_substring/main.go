// https://leetcode.com/problems/longest-palindromic-substring/
package main

func longestPalindrome(s string) string {
	from, to := 0, 0
	for i := 1; i < len(s); i++ {
		for _, p := range [][]int{{i, i}, {i - 1, i}} {
			l, r := p[0], p[1]
			for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			}
			l, r = l+1, r-1
			if r-l > to-from {
				from, to = l, r
			}
		}
	}
	return s[from:to+1]
}

func main() {}
