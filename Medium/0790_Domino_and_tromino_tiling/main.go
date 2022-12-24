// https://leetcode.com/problems/domino-and-tromino-tiling/
package main

func numTilings(n int) int {
	dom, trom, i, MOD := make([]int, n+2), make([]int, n+2), 3, int(1e9+7)
	for dom[0], dom[1], dom[2], trom[2] = 1, 1, 2, 1; i <= n; i++ {
		dom[i] = (dom[i-1] + dom[i-2] + trom[i-1]*2) % MOD
		trom[i] = (dom[i-2] + trom[i-1]) % MOD
	}
	return dom[n] % MOD
}

func main() {}
