// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package main

func maxProfit(prices []int) (ret int) {
	for i, stack := 0, []int{}; i < len(prices); i++ {
		for n := len(stack) - 1; n >= 0 && prices[i] <= stack[n]; n-- {
			stack = stack[:n]
		}
		stack = append(stack, prices[i])
		if diff := stack[len(stack)-1] - stack[0]; ret < diff {
			ret = diff
		}
	}
	return ret
}

func main() {}
