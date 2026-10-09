// https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/
package main

func minInsertions(s string) (res int) {
	var (
		left  int
		right bool
	)
	popPair := func() (popped bool) {
		if right {
			right, popped = false, true
			if left > 0 {
				left--
			} else {
				res++
			}
		}
		return popped
	}
	for _, ch := range s {
		switch ch {
		case '(':
			if popPair() {
				res++
			}
			left++
		case ')':
			if !popPair() {
				right = true
			}
		}
	}
	if popPair() {
		res++
	}
	return res + left*2
}

func main() {}
