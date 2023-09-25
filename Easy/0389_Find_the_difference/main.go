// https://leetcode.com/problems/find-the-difference/
package main

func findTheDifference(s string, t string) (ret byte) {
	cache := make([]int, 26)
	for i := range t {
		cache[t[i]-'a']++
	}
	for i := range s {
		cache[s[i]-'a']--
	}
	for i := range cache {
		if cache[i] > 0 {
			ret = byte(i + 'a')
		}
	}
	return ret
}
