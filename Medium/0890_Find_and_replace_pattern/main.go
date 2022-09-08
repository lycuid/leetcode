// https://leetcode.com/problems/find-and-replace-pattern/
package main

func findAndReplacePattern(words []string, pattern string) (ret []string) {
	var maps [26]byte
	var used [26]bool
NEXT_WORD:
	for _, word := range words {
		if len(word) == len(pattern) {
			for i := 0; i < 26; i++ {
				maps[i], used[i] = 27, false
			}
			for i := 0; i < len(word); i++ {
				ch1, ch2 := word[i]-'a', pattern[i]-'a'
				if maps[ch2] == 27 && !used[ch1] {
					maps[ch2], used[ch1] = ch1, true
				}
				if maps[ch2] != ch1 {
					continue NEXT_WORD
				}
			}
			ret = append(ret, word)
		}
	}
	return ret
}

func main() {}
