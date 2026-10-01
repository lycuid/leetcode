// https://leetcode.com/problems/valid-parentheses/
package main

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		switch ch := s[i]; ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			n := len(stack)
			if n == 0 {
				return false
			}
			if (ch == ')' && stack[n-1] != '(') ||
				(ch == ']' && stack[n-1] != '[') ||
				(ch == '}' && stack[n-1] != '{') {
				return false
			}
			stack = stack[:n-1]
		}
	}
	return len(stack) == 0
}

func main() {}
