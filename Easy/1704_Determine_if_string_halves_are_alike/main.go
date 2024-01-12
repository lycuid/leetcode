// https://leetcode.com/problems/determine-if-string-halves-are-alike/
package main

func IsVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func halvesAreAlike(s string) bool {
	var count int
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if IsVowel(s[i]) {
			count++
		}
		if IsVowel(s[j]) {
			count--
		}
	}
	return count == 0
}

func main() {}
