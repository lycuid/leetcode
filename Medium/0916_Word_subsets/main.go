// https://leetcode.com/problems/word-subsets/
package main

func wordSubsets(words1 []string, words2 []string) (ans []string) {
	var freq [26]int
	for _, word := range words2 {
		var freq2 [26]int
		for _, ch := range word {
			freq2[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if freq2[i] > freq[i] {
				freq[i] = freq2[i]
			}
		}
	}
NEXT_WORD:
	for _, word := range words1 {
		var freq2 [26]int
		for _, ch := range word {
			freq2[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if freq2[i] < freq[i] {
				continue NEXT_WORD
			}
		}
		ans = append(ans, word)
	}
	return ans
}

func main() {}
