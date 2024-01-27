// https://leetcode.com/problems/k-inverse-pairs-array/
package main

func kInversePairs(n int, k int) int {
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, k+1)
	}
	cache[0][0] = 1
	for i := 1; i <= n; i++ {
		for j, tmp := 0, 0; j <= k; j++ {
			if tmp += cache[i-1][j]; j >= i {
				tmp -= cache[i-1][j-i]
			}
			cache[i][j] = tmp % (1e9 + 7)
		}
	}
	return cache[n][k]
}

func main() {}
