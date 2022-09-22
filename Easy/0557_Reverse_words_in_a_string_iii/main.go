// https://leetcode.com/problems/reverse-words-in-a-string-iii/
package main

func reverseWords(s string) string {
	str := []byte(s)
	for start, end := 0, 0; end < len(str); start, end = end+1, end+1 {
		for ; end < len(str) && str[end] != ' '; end++ {
		}
		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			str[i], str[j] = str[j], str[i]
		}
	}
	return string(str)
}

func main() {}
