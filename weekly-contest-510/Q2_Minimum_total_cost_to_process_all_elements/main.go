// https://leetcode.com/contest/weekly-contest-510/problems/minimum-total-cost-to-process-all-elements/
package main

func minimumCost(nums []int, k int) (res int) {
	const mod = 1e9 + 7
	curr, cost := 0, k
	for _, num := range nums {
		var mul int
		if num > cost {
			mul = ((num - cost) + k - 1) / k
		}
		cost = cost + k*mul - num
		if mul > 0 {
			a, b := curr+1, curr+mul
			res += (((b - a + 1) % mod) * ((a + b) % mod) / 2) % mod
		}
		curr += mul
	}
	return res % mod
}

func main() {}
