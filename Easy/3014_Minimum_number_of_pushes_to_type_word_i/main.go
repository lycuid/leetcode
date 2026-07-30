// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/description/
package main

import "sort"

func minimumPushes(word string) (res int) {
	var cache [26]int
	for _, ch := range word {
		cache[ch-'a']++
	}
	sort.Slice(cache[:], func(i, j int) bool {
		return cache[i] > cache[j]
	})
	for i := range cache {
		res += cache[i] * (i/8 + 1)
	}
	return res
}

func main() {}
