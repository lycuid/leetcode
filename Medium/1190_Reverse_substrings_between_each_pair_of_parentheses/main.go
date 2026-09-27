// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/
package main

func reverseParentheses(s string) string {
	var (
		res    = make([]byte, 0, len(s))
		portal = make([]int, len(s))
		stack  = make([]int, 0, len(s))
	)

	for i, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, i)
		case ')':
			n := len(stack)
			portal[stack[n-1]], portal[i] = i, stack[n-1]
			stack = stack[:n-1]
		}
	}

	for i, inc := 0, 1; i < len(s); i += inc {
		switch ch := s[i]; ch {
		case '(', ')':
			i, inc = portal[i], -inc
		default:
			res = append(res, s[i])
		}
	}
	return string(res)
}

func main() {}
