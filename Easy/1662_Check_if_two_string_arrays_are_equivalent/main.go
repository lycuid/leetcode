// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var i, i1, j, j1 int
	for i < len(word1) && j < len(word2) {
		for ; i1 < len(word1[i]) && j1 < len(word2[j]); i1, j1 = i1+1, j1+1 {
			if word1[i][i1] != word2[j][j1] {
				return false
			}
		}
		if i1 == len(word1[i]) {
			if i++; i < len(word1) {
				i1 = 0
			}
		}
		if j1 == len(word2[j]) {
			if j++; j < len(word2) {
				j1 = 0
			}
		}
	}
	return i == len(word1) && i1 == len(word1[i-1]) && j == len(word2) && j1 == len(word2[j-1])
}

func main() {}
