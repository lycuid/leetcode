// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
package main

import "math"

func concatenatedBinary(n int) int {
	ret := 1
	for i := 2.; i <= float64(n); i++ {
		ret <<= int(math.Floor(math.Log2(i)) + 1)
		ret |= int(i)
		ret %= (1e9 + 7)
	}
	return ret
}

func main() {}
