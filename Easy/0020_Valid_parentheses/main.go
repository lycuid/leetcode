// https://leetcode.com/problems/valid-parentheses/
package main

func isValid(s string) bool {
	stack, cursor := make([]rune, len(s)), 0
	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			stack[cursor], cursor = ch, cursor+1
		} else if cursor > 0 && ((ch == '}' && stack[cursor-1] == '{') ||
			(ch == ']' && stack[cursor-1] == '[') ||
			(ch == ')' && stack[cursor-1] == '(')) {
			cursor--
		} else {
			return false
		}
	}
	return cursor == 0
}

func main() {}
