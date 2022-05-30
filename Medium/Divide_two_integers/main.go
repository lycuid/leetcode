// https://leetcode.com/problems/divide-two-integers/
package main

import "math"

func SignOf(x int) int {
	if Abs(x) != x {
		return -1
	}
	return 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinMax(x, min, max int) int {
	if min < x {
		x = min
	}
	if x < max {
		x = max
	}
	return x

}

func divide(dividend int, divisor int) (res int) {
	sign := SignOf(dividend) * SignOf(divisor)
	dd, dv := Abs(dividend), Abs(divisor)
	for dd >= dv {
		lshift := 0
		for dd > (dv << (lshift + 1)) {
			lshift++
		}
		dd -= (dv << lshift)
		res += 1 << lshift
	}
	res = MinMax(sign*res, math.MaxInt32, math.MinInt32)
	return res
}

func main() {}
