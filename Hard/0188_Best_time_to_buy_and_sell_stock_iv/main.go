// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
package main

func maxProfit(k int, prices []int) int {
	cost, profit := make([]int, k+1), make([]int, k+1)
	for i := range cost {
		cost[i] = 1e8 + 1
	}
	for _, price := range prices {
		for i := 1; i <= k; i++ {
			if c := price - profit[i-1]; c < cost[i] {
				cost[i] = c
			}
			if p := price - cost[i]; profit[i] < p {
				profit[i] = p
			}
		}
	}
	return profit[k]
}

func main() {}
