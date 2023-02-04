// https://leetcode.com/problems/permutation-in-string/
package main

func checkInclusion(s1 string, s2 string) bool {
	var f1, f2 [26]int
	if n, m := len(s1), len(s2); n <= m {
		eq := func() bool {
			for i := 0; i < 26; i++ {
				if f1[i] != f2[i] {
					return false
				}
			}
			return true
		}
		for i, ch := range s1 {
			f1[ch-'a'], f2[s2[i]-'a'] = f1[ch-'a']+1, f2[s2[i]-'a']+1
		}
		if eq() {
			return true
		}
		for i := 0; i+n < m; i++ {
			f2[s2[i]-'a']--
			if f2[s2[i+n]-'a']++; eq() {
				return true
			}
		}
	}
	return false
}

func main() {}
