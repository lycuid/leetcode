// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/
package main

func minLength(s string) int {
	var stack []byte
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == 'B' || ch == 'D' {
			if n := len(stack); n > 0 && stack[n-1] == ch-1 {
				stack = stack[:n-1]
				continue
			}
		}
		stack = append(stack, ch)
	}
	return len(stack)
}

func main() {}
