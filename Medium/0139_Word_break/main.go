// https://leetcode.com/problems/word-break/
package main

func wordBreak(s string, wordDict []string) bool {
	n, cache := len(s), make([]bool, len(s)+1)
	cache[n] = true
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < len(wordDict) && !cache[i]; j++ {
			if l := len(wordDict[j]); i+l <= n {
				cache[i] = wordDict[j] == s[i:i+l] && cache[i+l]
			}
		}
	}
	return cache[0]
}

func main() {}
