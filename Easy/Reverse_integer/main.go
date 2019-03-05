// https://leetcode.com/problems/reverse-integer/

package main

import (
	"math"
)

var MAX_INT32_DIV_10 int = int(math.Pow(2, 31)-1) / 10
var MIN_INT32_DIV_10 int = int(math.Pow(-2, 31)) / 10

func reverse(n int) int {
	number := 0
	for n != 0 {
		pop := n % 10
		if number > MAX_INT32_DIV_10 || ((number == MAX_INT32_DIV_10) && pop > 7) {
			return 0
		}
		if number < MIN_INT32_DIV_10 || ((number == MIN_INT32_DIV_10) && pop < -8) {
			return 0
		}

		number = (number * 10) + pop
		n /= 10
	}

	return number
}

func main() {}
