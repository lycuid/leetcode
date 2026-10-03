// https://leetcode.com/problems/longest-valid-parentheses/
package main

func longestValidParentheses(s string) (res int) {
	var (
		stack = make([]int, 0, len(s))
		cache = make([]int, len(s))
	)
	for i := range s {
		cache[i] = i
		switch n := len(stack); s[i] {
		case '(':
			if (n == 0 && i > 0) || (n > 0 && stack[n-1] != i-1) {
				cache[i] = cache[i-1]
			}
			stack = append(stack, i)
		case ')':
			if n > 0 && s[stack[n-1]] == '(' {
				cache[i] = cache[stack[n-1]]
				res = max(res, i-cache[i]+1)
				stack = stack[:n-1]
			} else {
				stack = append(stack, i)
			}
		}
	}
	return res
}

func main() {}
