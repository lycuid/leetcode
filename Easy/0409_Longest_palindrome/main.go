// https://leetcode.com/problems/longest-palindrome/
package main

func longestPalindrome(s string) (count int) {
	var (
		cache [128]int
		odd   int
	)
	for i := range s {
		cache[s[i]]++
	}
	for i := range cache {
		count += cache[i] >> 1 << 1
		odd |= cache[i] & 1
	}
	return count + odd
}

func main() {}
