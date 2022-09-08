// https://leetcode.com/problems/k-inverse-pairs-array/
package main

const MOD = 1e9 + 7

func kInversePairs(n int, k int) int {
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, k+1)
	}
	cache[0][0] = 1
	for i := 1; i <= n; i++ {
		var tmp int
		for j := 0; j <= k; j++ {
			tmp += cache[i-1][j]
			if j >= i {
				tmp -= cache[i-1][j-i]
			}
			cache[i][j] = (cache[i][j] + tmp) % MOD
		}
	}
	return cache[n][k]
}

func main() {}
