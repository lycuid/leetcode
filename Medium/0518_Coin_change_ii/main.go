// https://leetcode.com/problems/coin-change-ii/
package main

import "sort"

func change(amount int, coins []int) int {
	sort.Ints(coins)
	cache, n := make([][]int, amount+1), len(coins)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	for i := range cache[0] {
		cache[0][i] = 1
	}
	for amt := range cache {
		for j := 1; j <= n; j++ {
			cache[amt][j] = cache[amt][j-1]
			if prev := amt - coins[j-1]; prev >= 0 {
				cache[amt][j] += cache[prev][j]
			}
		}
	}
	return cache[amount][n]
}

func main() {}
