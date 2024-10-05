// https://leetcode.com/problems/permutation-in-string/
package main

func checkInclusion(s1 string, s2 string) bool {
	if n := len(s1); n <= len(s2) {
		var freq [2][26]int
		for i := range s1 {
			if freq[0][s1[i]-'a']++; i < n-1 {
				freq[1][s2[i]-'a']++
			}
		}
		for i := n - 1; i < len(s2); i++ {
			if freq[1][s2[i]-'a']++; freq[0] == freq[1] {
				return true
			}
			freq[1][s2[i-n+1]-'a']--
		}
	}
	return false
}

func main() {}
