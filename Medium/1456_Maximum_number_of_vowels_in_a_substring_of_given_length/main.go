// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
package main

func IsVowel(ch byte) int {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return 1
	}
	return 0
}

func maxVowels(s string, k int) (ret int) {
	max, i, n := 0, 0, len(s)
	for ; i < k; i++ {
		if max += IsVowel(s[i]); max > ret {
			ret = max
		}
	}
	for ; i < n; i++ {
		if max += IsVowel(s[i]) - IsVowel(s[i-k]); max > ret {
			ret = max
		}
	}
	return ret
}

func main() {}
