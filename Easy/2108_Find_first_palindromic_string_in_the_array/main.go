// https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
package main

func firstPalindrome(words []string) string {
NEXT_WORD:
	for _, word := range words {
		for i, n := 0, len(word); i <= n/2; i++ {
			if word[i] != word[n-i-1] {
				continue NEXT_WORD
			}
		}
		return word
	}
	return ""
}

func main() {}
