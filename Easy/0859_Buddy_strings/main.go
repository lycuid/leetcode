// https://leetcode.com/problems/buddy-strings/
package main

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	diff := 0
	for i := range s {
		if s[i] != goal[i] {
			diff++
		}
	}
	var freq [2][26]int
	if diff > 2 {
		return false
	}
	for i := range s {
		freq[0][s[i]-'a']++
	}
	for i := range goal {
		freq[1][goal[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if freq[0][i] != freq[1][i] {
			return false
		}
	}
	if diff == 2 {
		return true
	}
	for i := 0; i < 26; i++ {
		if freq[0][i] >= 2 {
			return true
		}
	}
	return false
}

func main() {}
