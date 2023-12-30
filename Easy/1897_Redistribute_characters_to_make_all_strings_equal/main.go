// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
package main

func makeEqual(words []string) bool {
	var freq [26]int
	for _, word := range words {
		for _, ch := range word {
			freq[ch-'a']++
		}
	}
	for i := range freq {
		if freq[i]%len(words) != 0 {
			return false
		}
	}
	return true
}

func main() {}
