// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
package main

func ValidWindow(sub string, n int, freq1 map[string]int) bool {
	freq2 := make(map[string]int)
	for i := 0; i < len(sub); i += n {
		word := sub[i : i+n]
		if _, found := freq1[word]; !found {
			return false
		}
		freq2[word]++
	}
	for k, v := range freq2 {
		if freq1[k] != v {
			return false
		}
	}
	return true
}

func findSubstring(s string, words []string) (ret []int) {
	freq1, n := make(map[string]int), len(words[0])
	for _, word := range words {
		freq1[word]++
	}
	window := len(words) * n
	for i := 0; i <= len(s)-window; i++ {
		if ValidWindow(s[i:i+window], n, freq1) {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {}
