// https://leetcode.com/problems/word-subsets/
package main

func wordSubsets(words1 []string, words2 []string) (ret []string) {
	var cache [26]int
	for _, word := range words2 {
		var freq [26]int
		for i := range word {
			freq[word[i]-'a']++
		}
		for i := range freq {
			cache[i] = max(cache[i], freq[i])
		}
	}

mainloop:
	for _, word := range words1 {
		var freq [26]int
		for i := range word {
			freq[word[i]-'a']++
		}
		for i := range freq {
			if freq[i] < cache[i] {
				continue mainloop
			}
		}
		ret = append(ret, word)
	}

	return ret
}

func main() {}
