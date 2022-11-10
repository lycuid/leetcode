// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
package main

func removeDuplicates(s string) string {
	stack := []byte{0}
	for i := range s {
		if n := len(stack) - 1; stack[n] == s[i] {
			stack = stack[:n]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack[1:])
}

func main() {}
