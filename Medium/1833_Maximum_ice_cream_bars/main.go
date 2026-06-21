// https://leetcode.com/problems/maximum-ice-cream-bars/
package main

func maxIceCream(costs []int, coins int) (res int) {
	var maxc int
	for _, cost := range costs {
		maxc = max(maxc, cost)
	}
	count := make([]int, maxc+1)
	for _, cost := range costs {
		count[cost]++
	}
	for i := 1; i < len(count) && coins/i > 0; i++ {
		n := min(coins/i, count[i])
		coins, res = coins-n*i, res+n
	}
	return res
}

func main() {}
