// https://leetcode.com/problems/sort-vowels-in-a-string/
package main

import "sort"

func IsVowel(ch rune) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}

func sortVowels(s string) string {
	var vowels, ret []rune
	for _, ch := range s {
		if IsVowel(ch) {
			vowels = append(vowels, ch)
		}
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})
	for _, ch := range s {
		if IsVowel(ch) {
			ret, vowels = append(ret, vowels[0]), vowels[1:]
		} else {
			ret = append(ret, ch)
		}
	}
	return string(ret)
}

func main() {}
