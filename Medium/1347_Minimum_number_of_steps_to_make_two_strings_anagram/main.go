// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
package main

func minSteps(s string, t string) (count int) {
	var freq [26]int
	for _, ch := range t {
		freq[ch-'a']++
	}
	for _, ch := range s {
		if freq[ch-'a'] > 0 {
			freq[ch-'a']--
		}
	}
	for i := range freq {
		count += freq[i]
	}
	return count
}

func main() {}
