// https://leetcode.com/problems/find-the-pivot-integer/
package main

import "math"

func pivotInteger(n int) int {
	a := math.Sqrt(float64((n * (n + 1)) / 2))
	if a_i := int(a); a == float64(a_i) {
		return a_i
	}
	return -1
}

func main() {}
