// https://leetcode.com/problems/smallest-palindromic-rearrangement-i/
package main

func smallestPalindrome(s string) string {
	var (
		cache [26]int
		res   = make([]byte, len(s))
		mid   = len(s) / 2
	)
	for _, ch := range s[:mid] {
		cache[ch-'a'] += 2
	}
	for i, j := 0, 0; i < len(cache); i++ {
		ch := byte(i) + 'a'
		for ; cache[i] > 0; cache[i], j = cache[i]-2, j+1 {
			res[j], res[len(res)-1-j] = ch, ch
		}
	}
	if len(s)%2 == 1 {
		res[mid] = s[mid]
	}
	return string(res)
}

func main() {}
