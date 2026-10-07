// https://leetcode.com/problems/remove-invalid-parentheses/
package main

func balance(s string) (int, int) {
	var left, right int
	for _, ch := range s {
		switch ch {
		case '(':
			left++
		case ')':
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left, right
}

func solve(prefix, s string, left, right int, res *[]string) {
	if left == 0 && right == 0 {
		ss := prefix + s
		if l, r := balance(ss); l == 0 && r == 0 {
			*res = append(*res, ss)
		}
		return
	}
	for i := range s {
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		if s[i] == '(' && left > 0 {
			solve(prefix+s[:i], s[i+1:], left-1, right, res)
		} else if s[i] == ')' && right > 0 {
			solve(prefix+s[:i], s[i+1:], left, right-1, res)
		}
	}
}

func removeInvalidParentheses(s string) (res []string) {
	left, right := balance(s)
	solve("", s, left, right, &res)
	return res
}

func main() {}
