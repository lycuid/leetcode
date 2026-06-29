// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
package main

func numOfStrings(patterns []string, word string) (count int) {
	for _, pat := range patterns {
		for i := 0; i <= len(word)-len(pat); i++ {
			if word[i:i+len(pat)] == pat {
				count++
				break
			}
		}
	}
	return count
}

func main() {}
