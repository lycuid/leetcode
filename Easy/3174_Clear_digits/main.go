// https://leetcode.com/problems/clear-digits/
package main

func clearDigits(s string) string {
	var stack []byte
	for i := range s {
		if ch := s[i]; ch >= '0' && ch <= '9' {
			if n := len(stack); n > 0 {
				stack = stack[:n-1]
			}
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

func main() {}
