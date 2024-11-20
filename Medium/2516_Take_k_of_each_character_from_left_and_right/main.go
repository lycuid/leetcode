// https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/
package main

func takeCharacters(s string, k int) (count int) {
	cache := [3]int{-k, -k, -k}
	for i := range s {
		cache[s[i]-'a']++
	}
	if cache[0] < 0 || cache[1] < 0 || cache[2] < 0 {
		return -1
	}
	cond := func(curr [3]int) bool {
		return curr[0] <= cache[0] && curr[1] <= cache[1] && curr[2] <= cache[2]
	}
	var curr [3]int
	for i, j := 0, 0; j < len(s); j++ {
		for curr[s[j]-'a']++; i <= j && !cond(curr); i++ {
			curr[s[i]-'a']--
		}
		count = max(count, j-i+1)
	}
	return len(s) - count
}

func main() {}
