// https://leetcode.com/problems/valid-word/
package main

import "strings"

func isValid(word string) bool {
	var hasVowel, hasConsonant bool
	if n := len(word); n >= 3 {
		word = strings.ToLower(word)
		for i := 0; i < n; i++ {
			ch := word[i]
			if !(ch >= '0' && ch <= '9') && !(ch >= 'a' && ch <= 'z') {
				return false
			}
			isVowel := ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
			hasVowel = hasVowel || isVowel
			hasConsonant = hasConsonant || (ch >= 'a' && ch <= 'z' && !isVowel)
		}
	}
	return hasVowel && hasConsonant
}

func main() {}
