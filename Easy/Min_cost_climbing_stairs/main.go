// https://leetcode.com/problems/min-cost-climbing-stairs/
package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	for i := 2; i < n; i++ {
		cost[i] = cost[i] + Min(cost[i-1], cost[i-2])
	}
	return Min(cost[n-1], cost[n-2])
}

func main() {}
