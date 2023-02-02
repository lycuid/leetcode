// https://leetcode.com/problems/verifying-an-alien-dictionary/
package main

func strcmp(word1, word2 string, chars [26]int) bool {
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) <= len(word2)
	}
	if chars[word1[0]-'a'] == chars[word2[0]-'a'] {
		return strcmp(word1[1:], word2[1:], chars)
	}
	return chars[word1[0]-'a'] < chars[word2[0]-'a']
}

func isAlienSorted(words []string, order string) bool {
	var chars [26]int
	for i, ch := range order {
		chars[ch-'a'] = i
	}
	for i := 0; i < len(words)-1; i++ {
		if !strcmp(words[i], words[i+1], chars) {
			return false
		}
	}
	return true
}

func main() {}
