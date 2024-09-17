// https://leetcode.com/problems/uncommon-words-from-two-sentences/
package main

func GetWord(s string) (string, string) {
	var i int
	for i < len(s) && s[i] != ' ' {
		i++
	}
	if i == len(s) {
		return s[:i], ""
	}
	return s[:i], s[i+1:]
}

func uncommonFromSentences(s1 string, s2 string) (words []string) {
	d1, d2 := make(map[string]int), make(map[string]int)
	for word := ""; len(s1) > 0; {
		word, s1 = GetWord(s1)
		d1[word]++
	}
	for word := ""; len(s2) > 0; {
		word, s2 = GetWord(s2)
		d2[word]++
	}
	for word, count := range d1 {
		if count == 1 {
			if _, found := d2[word]; !found {
				words = append(words, word)
			}
		}
	}
	for word, count := range d2 {
		if count == 1 {
			if _, found := d1[word]; !found {
				words = append(words, word)
			}
		}
	}
	return words
}

func main() {}
