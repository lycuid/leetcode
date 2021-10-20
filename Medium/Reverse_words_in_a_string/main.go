// https://leetcode.com/problems/reverse-words-in-a-string/
package main

import "strings"

func reverseWords(s string) string {
	ss := []string{}

	str := []byte{}
	for i, val := range s {
		if val == ' ' {
			if len(str) > 0 {
				ss = append(ss, string(str))
				str = []byte{}
			}
		} else {
			str = append(str, s[i])
		}
	}
	if len(str) > 0 {
		ss = append(ss, string(str))
	}

	l := len(ss) - 1
	for i := 0; i <= l/2; i++ {
		temp := ss[i]
		ss[i] = ss[l-i]
		ss[l-i] = temp
	}

	return strings.Join(ss, " ")
}

func main() {}
