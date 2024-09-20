// https://leetcode.com/problems/shortest-palindrome/
package main

func shortestPalindrome(s string) string {
	IsPalindrome := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	Reverse := func(s string) string {
		str := make([]byte, len(s))
		for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
			str[i], str[j] = s[j], s[i]
		}
		return string(str)
	}
	for i := len(s); i >= 0; i-- {
		if IsPalindrome(s[:i]) {
			return Reverse(s[i:]) + s
		}
	}
	return ""
}

func main() {}
