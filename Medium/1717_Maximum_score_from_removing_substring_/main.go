// https://leetcode.com/problems/maximum-score-from-removing-substrings/
package main

func maximumGain(s string, x int, y int) (score int) {
	ch := [2]byte{'a', 'b'}
	if y > x {
		x, y, ch[0], ch[1] = y, x, ch[1], ch[0]
	}
	var stack []byte
	for i := range s {
		if n := len(stack) - 1; n >= 0 && stack[n] == ch[0] && s[i] == ch[1] {
			stack, score = stack[:n], score+x
		} else {
			stack = append(stack, s[i])
		}
	}
	ch[0], ch[1] = ch[1], ch[0]
	s, stack, x, y = string(stack), []byte{}, y, x
	for i := range s {
		if n := len(stack) - 1; n >= 0 && stack[n] == ch[0] && s[i] == ch[1] {
			stack, score = stack[:n], score+x
		} else {
			stack = append(stack, s[i])
		}
	}
	return score
}

func main() {}
