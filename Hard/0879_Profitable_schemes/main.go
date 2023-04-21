// https://leetcode.com/problems/profitable-schemes/
package main

const MOD = 1e9 + 7

func profitableSchemes(n int, minProfit int, group []int, profit []int) (ret int) {
	cache := make([][][]int, len(group)+1)
	for i := range cache {
		cache[i] = make([][]int, n+1)
		for j := range cache[i] {
			cache[i][j] = make([]int, minProfit+1)
		}
	}
	cache[0][0][0] = 1
	for i := 1; i <= len(group); i++ {
		g, p := group[i-1], profit[i-1]
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if cache[i][j][k] = cache[i-1][j][k]; j >= g {
					k_ := k - p
					if k_ < 0 {
						k_ = 0
					}
					cache[i][j][k] = (cache[i][j][k] + cache[i-1][j-g][k_]) % MOD
				}
			}
		}
	}
	for i := 0; i <= n; i++ {
		ret = (ret + cache[len(group)][i][minProfit]) % MOD
	}
	return ret
}

func main() {}
