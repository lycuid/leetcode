// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

func lengthOfLongestSubstring(s string) (ret int) {
	var cache [129]bool
	for i, j := 0, 0; i < len(s); i++ {
		for ; cache[s[i]] && j <= i; j++ {
			cache[s[j]] = false
		}
		if i-j+1 > ret {
			ret = i - j + 1
		}
		cache[s[i]] = true
	}
	return ret
}

func main() {}
