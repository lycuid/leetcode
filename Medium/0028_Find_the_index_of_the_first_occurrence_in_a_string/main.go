// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
package main

func strStr(haystack string, needle string) int {
	for i, j, hl, nl := 0, 0, len(haystack), len(needle); i <= hl-nl; i++ {
		for j = i; j < hl && j-i < nl && haystack[j] == needle[j-i]; j++ {
		}
		if j-i == nl {
			return i
		}
	}
	return -1
}

func main() {}
