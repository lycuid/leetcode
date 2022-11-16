// https://leetcode.com/problems/guess-number-higher-or-lower/
package main

func guess(num int) int

func guessNumber(n int) (mid int) {
	for l, r := 1, n; ; {
		mid = (l + r) / 2
		if p := guess(mid); p == 1 {
			l = mid + 1
		} else if p == -1 {
			r = mid
		} else {
			break
		}
	}
	return mid
}

func main() {}
