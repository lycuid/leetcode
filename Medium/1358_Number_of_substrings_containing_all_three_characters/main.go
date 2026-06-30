// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
package main

func numberOfSubstrings(s string) (res int) {
	cache := [3]int{-1, -1, -1}
	for i, ch := range s {
		cache[ch-'a'] = i
		res += 1 + min(cache[0], cache[1], cache[2])
	}
	return res
}

func main() {}
