// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/
package main

func Reverse(chars []byte) []byte {
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return chars
}

func Aux(chars []byte) []byte {
	for i := 0; i < len(chars); i++ {
		if chars[i] == '(' {
			j := i + 1
			for depth := 1; j < len(chars); j++ {
				switch chars[j] {
				case '(':
					depth++
				case ')':
					depth--
				}
				if depth == 0 {
					break
				}
			}
			return append(chars[:i], append(Reverse(Aux(chars[i+1:j])), Aux(chars[j+1:])...)...)
		}
	}
	return chars
}

func reverseParentheses(s string) string {
	return string(Aux([]byte(s)))
}

func main() {}
