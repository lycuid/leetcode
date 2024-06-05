// https://leetcode.com/problems/find-common-characters/
package main

import (
	"fmt"
	"math"
)

func commonChars(words []string) (ret []string) {
	var cache [26]int
	for i := range cache {
		cache[i] = math.MaxInt
	}
	for _, word := range words {
		var freq [26]int
		for i := 0; i < len(word); i++ {
			freq[word[i]-'a']++
		}
		for i := range freq {
			if freq[i] < cache[i] {
				cache[i] = freq[i]
			}
		}
	}
	for i, count := range cache {
		if count < math.MaxInt {
			for ; count > 0; count-- {
				ret = append(ret, fmt.Sprintf("%c", i+'a'))
			}
		}
	}
	return ret
}

func main() {}
