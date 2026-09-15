// https://leetcode.com/problems/maximum-number-of-non-overlapping-palindrome-substrings/
package main

import "sort"

func maxPalindromes(s string, k int) (res int) {
	cache := make([][2]int, 1, len(s))
	cache[0] = [2]int{-1, -1}
	for i := range s {
		for l, r := i, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l+1 >= k {
				cache = append(cache, [2]int{l, r})
				break
			}
		}
		for l, r := i-1, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l+1 >= k {
				cache = append(cache, [2]int{l, r})
				break
			}
		}
	}
	sort.Slice(cache, func(i, j int) bool {
		if cache[i][1] == cache[j][1] {
			return cache[i][0] < cache[j][0]
		}
		return cache[i][1] < cache[j][1]
	})
	for i := 1; i < len(cache); i++ {
		if cache[i][0] > cache[0][1] {
			cache[0] = cache[i]
			res++
		}
	}
	return res
}

func main() {}
