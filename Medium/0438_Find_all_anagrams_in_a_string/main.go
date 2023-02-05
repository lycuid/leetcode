// https://leetcode.com/problems/find-all-anagrams-in-a-string/
package main

func findAnagrams(s string, p string) (ret []int) {
	var f1, f2 [26]int
	if n, m := len(s), len(p); n >= m {
		eq := func() bool {
			for i := 0; i < 26; i++ {
				if f1[i] != f2[i] {
					return false
				}
			}
			return true
		}
		for i := range p {
			f1[p[i]-'a'], f2[s[i]-'a'] = f1[p[i]-'a']+1, f2[s[i]-'a']+1
		}
		for i := 0; i+m <= n; i++ {
			if eq() {
				ret = append(ret, i)
			}
			if f2[s[i]-'a']--; i+m < n {
				f2[s[i+m]-'a']++
			}
		}
	}
	return ret
}

func main() {}
