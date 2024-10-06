// https://leetcode.com/problems/sentence-similarity-iii/
package main

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words := [2][]string{
		strings.Split(sentence1, " "),
		strings.Split(sentence2, " "),
	}
	if len(words[0]) > len(words[1]) {
		words[0], words[1] = words[1], words[0]
	}
	Prefix := func(s1 []string) []string {
		for i := 0; len(s1) > 0; i++ {
			if s1[0] != words[1][i] {
				break
			}
			s1 = s1[1:]
		}
		return s1
	}
	Suffix := func(s1 []string) []string {
		for i := len(words[1]) - 1; len(s1) > 0 && i >= 0; i-- {
			n := len(s1) - 1
			if s1[n] != words[1][i] {
				break
			}
			s1 = s1[:n]
		}
		return s1
	}
	return len(Suffix(Prefix(words[0]))) == 0 || len(Prefix(Suffix(words[0]))) == 0
}

func main() {}
