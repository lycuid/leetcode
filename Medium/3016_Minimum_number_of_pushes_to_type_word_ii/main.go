// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/
package main

import "sort"

func minimumPushes(word string) (ret int) {
	var cache [26]int
	for i := range word {
		cache[word[i]-'a']++
	}
	sort.Slice(cache[:], func(i, j int) bool { return cache[i] > cache[j] })
	for i, j := 1, 0; i*8 < len(cache)+8; i++ {
		for ; j < i*8 && j < len(cache); j++ {
			ret += cache[j] * i
		}
	}
	return ret
}

func main() {}
