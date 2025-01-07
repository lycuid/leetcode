// https://leetcode.com/problems/string-matching-in-an-array/
package main

func isSubString(s1, s2 string) bool {
	if n, m := len(s1), len(s2); n <= m {
		for i := 0; i <= m-n; i++ {
			if s1 == s2[i:i+n] {
				return true
			}
		}
	}
	return false
}

func stringMatching(words []string) (ret []string) {
	for i, word := range words {
		for j := range words {
			if i != j {
				if isSubString(word, words[j]) {
					ret = append(ret, word)
					break
				}
			}
		}
	}
	return ret
}

func main() {}
