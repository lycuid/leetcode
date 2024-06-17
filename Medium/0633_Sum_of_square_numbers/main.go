// https://leetcode.com/problems/sum-of-square-numbers/
package main

import "math"

func judgeSquareSum(c int) bool {
	r, num := int64(math.Ceil(math.Sqrt(float64(c)))), int64(c)
	for l := int64(0); l <= r; {
		if sum := l*l + r*r; sum < num {
			l++
		} else if sum > num {
			r--
		} else {
			return true
		}
	}
	return false
}

func main() {}
