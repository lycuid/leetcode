// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
package main

func removeOccurrences(s string, part string) string {
	stack, n := make([]byte, 0, len(s)), len(part)
	for i := range s {
		stack = append(stack, s[i])
		for len(stack) >= n && string(stack[len(stack)-n:]) == part {
			stack = stack[:len(stack)-n]
		}
	}
	return string(stack)
}

func main() {}
