// https://leetcode.com/problems/buy-two-chocolates/
package main

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if price := money - prices[0] - prices[1]; price >= 0 {
		return price
	}
	return money
}

func main() {}
