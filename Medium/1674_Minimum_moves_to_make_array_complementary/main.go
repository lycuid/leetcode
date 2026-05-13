// https://leetcode.com/problems/minimum-moves-to-make-array-complementary/
package main

import "sort"

func minMoves(nums []int, limit int) int {
	var (
		n     = len(nums)
		res   = 2*n + 1
		cache = make([]int, 2*limit+1)
		low   = make([]int, 0, n)
		high  = make([]int, 0, n)
	)

	for i := 0; i < n>>1; i++ {
		a, b := nums[i], nums[n-1-i]
		cache[a+b]++
		low = append(low, min(a+1, b+1))
		high = append(high, max(a+limit, b+limit))
	}
	sort.Ints(low)
	sort.Ints(high)

	find := func(num int) int {
		if num <= limit {
			return sort.Search(len(low), func(i int) bool {
				return low[i] > num
			})
		}
		return len(high) - sort.Search(len(high), func(i int) bool {
			return high[i] >= num
		})
	}

	for i, zero := range cache {
		one := max(0, find(i)-zero)
		two := n>>1 - (zero + one)
		res = min(res, one+2*two)
	}
	return res
}

func main() {}
