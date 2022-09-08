// https://leetcode.com/problems/valid-palindrome-ii/
package main

func palindrome(str string, s, e, c int) bool {
	if c < 0 {
		return false
	}
	for ; s < e; s++ {
		if str[s] != str[e] {
			return palindrome(str, s+1, e, c-1) || palindrome(str, s, e-1, c-1)
		}
		e--
	}
	return true
}

func validPalindrome(s string) bool {
	return palindrome(s, 0, len(s)-1, 1)
}

func main() {}
