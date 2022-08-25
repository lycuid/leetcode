// https://leetcode.com/problems/ransom-note/
package main

func canConstruct(ransomNote string, magazine string) bool {
	var freq [26]int
	for _, ch := range magazine {
		freq[ch-'a']++
	}
	for _, ch := range ransomNote {
		if freq[ch-'a']--; freq[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {}
