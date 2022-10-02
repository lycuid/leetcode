// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
package main

func numRollsToTarget(n int, k int, target int) (ret int) {
	cache := make([][]int, target+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
		if 0 < i && i <= k {
			cache[i][1] = 1
		}
	}
	for i := 1; i <= target; i++ {
		for j := 2; j <= n; j++ {
			for p := 1; p <= i && p <= k; p++ {
				cache[i][j] += cache[i-p][j-1]
				cache[i][j] %= (1e9 + 7)
			}
		}
	}
	return cache[target][n]
}
