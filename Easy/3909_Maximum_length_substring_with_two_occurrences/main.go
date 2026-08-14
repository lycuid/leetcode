// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/
package main

func maximumLengthSubstring(s string) (res int) {
	var cache [26]int
	for i, j := 0, 0; i < len(s); i++ {
		for cache[s[i]-'a']++; cache[s[i]-'a'] > 2; j++ {
			cache[s[j]-'a']--
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {}
