// https://leetcode.com/problems/valid-anagram/
package main

func isAnagram(s string, t string) bool {
	var seen [26]int
	for _, ch := range s {
		seen[ch - 'a']++
	}
	for _, ch := range t {
		seen[ch - 'a']--
		if seen[ch - 'a'] < 0 {
			return false
		}
	}
	sum := 0
	for _, n := range seen {
		sum += n
	}
	return sum == 0
}

func main() {}
