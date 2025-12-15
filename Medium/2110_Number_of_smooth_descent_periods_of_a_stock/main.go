// https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/
package main

import "math"

func getDescentPeriods(prices []int) (count int64) {
	prices = append([]int{math.MaxInt}, prices...)
	for i, j := 1, 1; i < len(prices); i++ {
		if prices[i]+1 != prices[i-1] {
			j = i
		}
		count += int64(i - j + 1)
	}
	return count
}

func main() {}
