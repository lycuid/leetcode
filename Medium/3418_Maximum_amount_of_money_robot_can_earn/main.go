// https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/
package main

import "math"

func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	cache := make([][][3]int, m+1)
	for i := range cache {
		cache[i] = make([][3]int, n+1)
		for j := range cache[i] {
			cache[i][j] = [3]int{math.MinInt32, math.MinInt32, math.MinInt32}
		}
	}
	cache[1][0][0] = 0
	for i := range coins {
		for j, val := range coins[i] {
			cache[i+1][j+1][0] = val + max(cache[i][j+1][0], cache[i+1][j][0])
			cache[i+1][j+1][1] = max(val+cache[i][j+1][1], cache[i][j+1][0], val+cache[i+1][j][1], cache[i+1][j][0])
			cache[i+1][j+1][2] = max(val+cache[i][j+1][2], cache[i][j+1][1], val+cache[i+1][j][2], cache[i+1][j][1])
		}
	}
	return max(cache[m][n][0], cache[m][n][1], cache[m][n][2])
}

func main() {}
