// https://leetcode.com/problems/valid-anagram/
package main

func isAnagram(s string, t string) bool {
	cache := make(map[rune]int)
	for _, ch := range s {
		cache[ch]++
	}
	for _, ch := range t {
		cache[ch]--
	}
	for i := range cache {
		if cache[i] != 0 {
			return false
		}
	}
	return true
}

func main() {}
