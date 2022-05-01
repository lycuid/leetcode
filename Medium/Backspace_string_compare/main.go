// https://leetcode.com/problems/backspace-string-compare/
package main

import "fmt"

func resolve(str []byte) (cursor int) {
	for i, ch := range str {
		if ch == '#' {
			if cursor > 0 {
				cursor--
			}
			continue
		}
		str[cursor] = str[i]
		cursor++
	}
	return cursor
}

func backspaceCompare(s string, t string) bool {
	str1, str2 := []byte(s), []byte(t)
	cursor1, cursor2 := resolve(str1), resolve(str2)
	if cursor1 != cursor2 {
		return false
	}
	for i := 0; i < cursor1; i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "c#d#"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
}
