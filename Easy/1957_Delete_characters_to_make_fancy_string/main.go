// https://leetcode.com/problems/delete-characters-to-make-fancy-string/
package main

func makeFancyString(s string) string {
	stack := make([]rune, 0, len(s))
	for _, ch := range s {
		stack = append(stack, ch)
		for n := len(stack); n >= 3 && stack[n-1] == stack[n-2] && stack[n-1] == stack[n-3]; n = len(stack) {
			stack = stack[:n-1]
		}
	}
	return string(stack)
}

func main() {}
