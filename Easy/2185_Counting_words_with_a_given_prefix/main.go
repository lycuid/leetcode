// https://leetcode.com/problems/counting-words-with-a-given-prefix/
package main

import "strings"

func prefixCount(words []string, pref string) (count int) {
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}
	return count
}

func main() {}
