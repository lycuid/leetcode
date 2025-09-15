// https://leetcode.com/problems/maximum-number-of-words-you-can-type/
package main

func canBeTypedWords(text string, brokenLetters string) (count int) {
	var broken, word uint32
	for i := range brokenLetters {
		broken |= 1 << (brokenLetters[i] - 'a')
	}
	for i := range text {
		switch ch := text[i]; ch {
		case ' ':
			if word&broken == 0 {
				count++
			}
			word = 0
		default:
			word |= 1 << (ch - 'a')
		}
	}
	if word&broken == 0 {
		count++
	}
	return count
}

func main() {}
