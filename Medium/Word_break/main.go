// https://leetcode.com/problems/word-break/
package main

func wordBreak(s string, wordDict []string) bool {
	cache, l := make([]bool, len(s)+1), len(s)
	cache[l] = true
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(wordDict) && !cache[i]; j++ {
			if wl := len(wordDict[j]); l-i >= wl && wordDict[j] == s[i:i+wl] {
				cache[i] = cache[i+wl]
			}
		}
	}
	return cache[0]
}

func main() {}
