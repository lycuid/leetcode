// https://leetcode.com/problems/coin-change/
package main

import "math"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt - 1
	}
	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = Min(dp[i], 1+dp[i-coin])
			}
		}
	}
	if dp[amount] < math.MaxInt-1 {
		return dp[amount]
	}
	return -1
}

func main() {}
