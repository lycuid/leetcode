// https://leetcode.com/problems/reverse-prefix-of-word/
package main

func reversePrefix(word string, ch byte) string {
	chars, i := []byte(word), 0
	for i < len(chars) && chars[i] != ch {
		i++
	}
	if i < len(chars) {
		for j := 0; j < i; i, j = i-1, j+1 {
			chars[i], chars[j] = chars[j], chars[i]
		}
	}
	return string(chars)
}

func main() {}
