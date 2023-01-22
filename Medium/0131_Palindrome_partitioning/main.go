// https://leetcode.com/problems/palindrome-partitioning/
package main

func IsPalindrome(s string, l, r int) bool {
	for ; l <= r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	cache, n := make([][][]string, len(s)+1), len(s)
	cache[0] = append(cache[0], []string{})
	for r := 0; r < n; r++ {
		for l, last := r, 0; l >= 0; l-- {
			if IsPalindrome(s, l, r) {
				for _, subs := range cache[l] {
					cache[r+1], last = append(cache[r+1], []string{}), len(cache[r+1])
					for _, sub := range subs {
						cache[r+1][last] = append(cache[r+1][last], sub)
					}
					cache[r+1][last] = append(cache[r+1][last], s[l:r+1])
				}
			}
		}
	}
	return cache[n]
}

func main() {}
