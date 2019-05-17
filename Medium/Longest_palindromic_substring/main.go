// https://leetcode.com/problems/longest-palindromic-substring/

package main

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < (l / 2); i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	palindrome := ""

	for i := 0; i < len(s)-len(palindrome); i++ {
		l := len(palindrome)
		for j := i + l; j < len(s); j++ {
			substr := s[i : j+1]
			if s[i] == s[j] && isPalindrome(substr) {
				palindrome = substr
			}
		}
	}
	return palindrome
}

func main() {}
