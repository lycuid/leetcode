// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
package main

func minRemoveToMakeValid(s string) string {
	chars := []byte(s)
	var stack []int
	for i, ch := range chars {
		switch ch {
		case '(':
			stack = append(stack, i)
			break
		case ')':
			if n := len(stack); n > 0 {
				stack = stack[:n-1]
			} else {
				chars[i] = '#'
			}
			break
		}
	}
	for len(stack) > 0 {
		chars[stack[0]], stack = '#', stack[1:]
	}
	var i int
	for j := 0; j < len(chars); j++ {
		if chars[j] != '#' {
			chars[i], i = chars[j], i+1
		}
	}
	return string(chars[:i])
}

func main() {}
