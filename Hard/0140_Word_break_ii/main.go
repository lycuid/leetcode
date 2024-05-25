// https://leetcode.com/problems/word-break-ii/
package main

import (
	"fmt"
	"strings"
)

func Aux(s string, start int, cache [][]string, dict map[string]bool) []string {
	for i := start; i < len(s); i++ {
		if _, found := dict[s[start:i+1]]; found {
			if cache[i+1] == nil {
				Aux(s, i+1, cache, dict)
			}
			for _, sentence := range cache[i+1] {
				cache[start] = append(cache[start], fmt.Sprintf("%s %s", s[start:i+1], sentence))
			}
		}
	}
	return cache[start]
}

func wordBreak(s string, wordDict []string) []string {
	var (
		cache = make([][]string, len(s)+1)
		dict  = make(map[string]bool)
	)
	cache[len(s)] = []string{""}
	for _, word := range wordDict {
		dict[word] = true
	}
	ret := Aux(s, 0, cache, dict)
	for i, val := range ret {
		ret[i] = strings.TrimRight(val, " ")
	}
	return ret
}

func main() {}
