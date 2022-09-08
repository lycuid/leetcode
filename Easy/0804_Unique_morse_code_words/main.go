// https://leetcode.com/problems/unique-morse-code-words/
package main

func uniqueMorseRepresentations(words []string) int {
	morse := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....",
		"..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.",
		"...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	unique := make(map[string]bool)
	for _, word := range words {
		var code string
		for _, ch := range word {
			code += morse[ch-'a']
		}
		unique[code] = true
	}
	return len(unique)
}

func main() {}
