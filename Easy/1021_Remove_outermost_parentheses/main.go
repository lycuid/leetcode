// https://leetcode.com/problems/remove-outermost-parentheses/
package main

func removeOuterParentheses(s string) string {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		j := i + 1
		for weight := 0; weight >= 0 && j < len(s); j++ {
			switch s[j] {
			case '(':
				weight++
			case ')':
				weight--
			}
		}
		for i++; i < j-1; i++ {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func main() {}
