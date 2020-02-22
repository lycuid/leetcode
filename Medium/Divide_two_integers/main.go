// https://leetcode.com/problems/divide-two-integers/

package main

import "math"

func negate(i int) int {
	return i - (i + i)
}

func multiply(i, j int) int {
	if j == 0 {
		return 0
	}

	if j > 0 {
		return i + multiply(i, j-1)
	}

	return negate(multiply(i, negate(j)))
}

func signOf(i int) int {
	if i < 0 {
		return -1
	}
	return 1
}

func abs(i int) int {
	if i < 0 {
		return negate(i)
	}
	return i
}

func divide(dividend int, divisor int) (res int) {
	isNegative := multiply(signOf(dividend), signOf(divisor)) < 0
	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend >= divisor {
		dividend -= divisor
		res++
	}

	if isNegative {
		res = negate(res)
		if res < math.MinInt32 {
			res = math.MinInt32
		}
	} else {
		if res > math.MaxInt32 {
			res = math.MaxInt32
		}
	}
	return res
}

func main() {}
