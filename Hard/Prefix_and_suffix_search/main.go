// https://leetcode.com/problems/prefix-and-suffix-search/
package main

import "fmt"

type WordFilter struct {
	inner map[string]int
}

func Constructor(words []string) WordFilter {
	inner := make(map[string]int)
	for index, word := range words {
		for i := len(word); i >= 0; i-- {
			for j := 0; j <= len(word); j++ {
				s := fmt.Sprintf("%s.%s", word[:i], word[j:])
				if inner[s] <= index {
					inner[s] = index
				}
			}
		}
	}
	return WordFilter{inner}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	s := fmt.Sprintf("%s.%s", prefix, suffix)
	if index, found := this.inner[s]; found {
		return index
	}
	return -1
}

func main() {}
