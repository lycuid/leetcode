// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
package main

func maxProfit(prices []int) int {
	cache, n := make([]int, len(prices)+2), len(prices)
	for i, j := n-1, 0; i >= 0; i-- {
		for j, cache[i] = i+1, cache[i+1]; j < n; j++ {
			if profit := prices[j] - prices[i] + cache[j+2]; profit > cache[i] {
				cache[i] = profit
			}
		}
	}
	return cache[0]
}

func main() {}
