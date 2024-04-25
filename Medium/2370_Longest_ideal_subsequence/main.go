// https://leetcode.com/problems/longest-ideal-subsequence/
package main

func longestIdealString(s string, k int) (ret int) {
	var cache [26]int
	for i := range s {
		ch := int(s[i] - 'a')
		curr, l, r := cache[ch], ch-k, ch+k
		if l < 0 {
			l = 0
		}
		if r > 25 {
			r = 25
		}
		for j := l; j <= r; j++ {
			if n := cache[j] + 1; n > curr {
				curr = n
			}
		}
		if cache[ch] = curr; cache[ch] > ret {
			ret = cache[ch]
		}
	}
	return ret
}

func main() {}
