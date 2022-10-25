// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	cache := ""
	for _, word := range word1 {
		cache += word
	}
	for _, word := range word2 {
		if n := len(word); len(cache) >= n && cache[:n] == word {
			cache = cache[n:]
		} else {
			return false
		}
	}
	return len(cache) == 0
}

func main() {}
