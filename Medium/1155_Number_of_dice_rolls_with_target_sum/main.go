// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
package main

func numRollsToTarget(n int, k int, target int) (ret int) {
	cache := make([][]int, target+1)
	for i := range cache {
		if cache[i] = make([]int, n+1); 0 < i && i <= k {
			cache[i][1] = 1
		}
	}
	for t := 1; t <= target; t++ {
		for j := 2; j <= n; j++ {
			for roll := 1; roll <= t && roll <= k; roll++ {
				cache[t][j] = (cache[t][j] + cache[t-roll][j-1]) % (1e9 + 7)
			}
		}
	}
	return cache[target][n]
}
