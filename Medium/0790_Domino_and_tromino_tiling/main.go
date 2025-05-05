// https://leetcode.com/problems/domino-and-tromino-tiling/
package main

func numTilings(n int) int {
	cache := make([]int, n+3)
	cache[0], cache[1], cache[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		cache[i] = (2*cache[i-1] + cache[i-3]) % (1e9 + 7)
	}
	return cache[n]
}

func main() {}
