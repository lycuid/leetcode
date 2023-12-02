// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
package main

func Freq(chars string) (freq [26]rune) {
	for _, ch := range chars {
		freq[ch-'a']++
	}
	return freq
}

func countCharacters(words []string, chars string) (count int) {
	f_chars := Freq(chars)
mainloop:
	for _, word := range words {
		freq := Freq(word)
		for i, f := range freq {
			if f > f_chars[i] {
				continue mainloop
			}
		}
		count += len(word)
	}
	return count
}

func main() {}
