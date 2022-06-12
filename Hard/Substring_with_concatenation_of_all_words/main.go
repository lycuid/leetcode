// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
package main

func ValidWindow(start, end, word_len int, s *string, word_freq *map[string]int) bool {
	window_freq := make(map[string]int)
	for i := start; i < end; i += word_len {
		word := (*s)[i : i+word_len]
		if _, found := (*word_freq)[word]; !found {
			return false
		}
		window_freq[word]++
	}
	for word, f1 := range window_freq {
		if f2, found := (*word_freq)[word]; !found || f1 != f2 {
			return false
		}
	}
	return true
}

func findSubstring(s string, words []string) (res []int) {
	word_len, word_freq := len(words[0]), make(map[string]int)
	window := len(words) * word_len
	for i := range words {
		word_freq[words[i]]++
	}
	for i := 0; i <= len(s)-window; i++ {
		if ValidWindow(i, i+window, word_len, &s, &word_freq) {
			res = append(res, i)
		}
	}
	return res
}

func main() {}
