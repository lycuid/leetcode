// https://leetcode.com/problems/maximum-score-from-removing-substrings/
package main

func maximumGain(s string, x int, y int) (score int) {
	stack, a, b := make([]rune, 0, len(s)), 'a', 'b'
	if y > x {
		a, b, x, y = 'b', 'a', y, x
	}
	for _, ch := range s {
		stack = append(stack, ch)
		for n := len(stack); n >= 2 && stack[n-2] == a && stack[n-1] == b; n = len(stack) {
			stack, score = stack[:n-2], score+x
		}
	}
	s, stack = string(stack), stack[:0]
	for _, ch := range s {
		stack = append(stack, ch)
		for n := len(stack); n >= 2 && stack[n-2] == b && stack[n-1] == a; n = len(stack) {
			stack, score = stack[:n-2], score+y
		}
	}
	return score
}

func main() {}
