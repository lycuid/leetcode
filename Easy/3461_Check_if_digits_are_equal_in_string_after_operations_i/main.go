// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/
package main

func hasSameDigits(s string) bool {
	if len(s) == 2 {
		return s[0] == s[1]
	}
	cache := make([]byte, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		cache[i] = ((s[i] - '0' + s[i+1] - '0') % 10) + '0'
	}
	return hasSameDigits(string(cache))
}

func main() {}
