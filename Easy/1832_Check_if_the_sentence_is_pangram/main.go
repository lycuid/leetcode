// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
package main

func checkIfPangram(sentence string) bool {
	var chars [26]bool
	for i, found := 0, 0; i < len(sentence); i++ {
		if ch := sentence[i]; !chars[ch-'a'] {
			if chars[ch-'a'], found = true, found+1; found == 26 {
				return true
			}
		}
	}
	return false
}

func main() {}
