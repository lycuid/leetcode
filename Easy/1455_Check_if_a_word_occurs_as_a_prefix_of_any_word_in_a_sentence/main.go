// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, word := range strings.Fields(sentence) {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}
	return -1
}

func main() {}
