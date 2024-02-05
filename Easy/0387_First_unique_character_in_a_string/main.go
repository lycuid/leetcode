// https://leetcode.com/problems/first-unique-character-in-a-string/
package main

func firstUniqChar(s string) int {
	var freq [26]int
	for _, ch := range s {
		freq[ch-'a']++
	}
	for i := range s {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {}
