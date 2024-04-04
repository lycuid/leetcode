// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
package main

func maxDepth(s string) (count int) {
	for i, max := 0, 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			if max++; max > count {
				count = max
			}
			break
		case ')':
			max--
			break
		}
	}
	return count
}

func main() {}
