// https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
package main

func smallestSubsequence(s string) string {
	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
	}
	st, mask := make([]byte, 0, len(s)), 0
	for i := range s {
		cnt[s[i]-'a']--
		bit := 1 << (s[i] - 'a')
		if mask&bit != 0 {
			continue
		}
		for n := len(st) - 1; n >= 0 && s[i] < st[n] && cnt[st[n]-'a'] > 0; n = len(st) - 1 {
			mask, st = mask&^(1<<(st[n]-'a')), st[:n]
		}
		mask, st = mask|bit, append(st, s[i])
	}
	return string(st)
}

func main() {}
