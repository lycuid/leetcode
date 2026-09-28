// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
package main

func maxDepth(s string) (res int) {
	var depth int
	for _, ch := range s {
		switch ch {
		case '(':
			depth++
			res = max(res, depth)
		case ')':
			depth--
		}
	}
	return res
}

func main() {}
