// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
package main

func numRollsToTarget(n int, k int, target int) (ret int) {
	cache := make([][]int, n+1)
	for i := range cache {
		if cache[i] = make([]int, target+1); i == 1 {
			for j := 0; j <= k && j <= target; j++ {
				cache[i][j] = 1
			}
		}
	}
	for i := 1; i <= n; i++ {
		for t := 2; t <= target; t++ {
			for roll := 1; roll <= k && roll < t; roll++ {
				cache[i][t] = (cache[i][t] + cache[i-1][t-roll]) % (1e9 + 7)
			}
		}
	}
	return cache[n][target]
}

func main() {}
