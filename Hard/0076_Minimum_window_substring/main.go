// https://leetcode.com/problems/minimum-window-substring/
package main

func Check(c1, c2 [129]int) bool {
	for i := 0; i < 129; i++ {
		if c1[i] < c2[i] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) (ret string) {
	var cache, chars [129]int
	for _, ch := range t {
		chars[ch]++
	}
	for i, j := 0, 0; i < len(s); i++ {
		for cache[s[i]]++; Check(cache, chars); j++ {
			if cache[s[j]]--; len(ret) == 0 || i+1-j < len(ret) {
				ret = s[j : i+1]
			}
		}
	}
	return ret
}

func main() {}
