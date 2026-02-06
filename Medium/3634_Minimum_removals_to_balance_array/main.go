// https://leetcode.com/problems/minimum-removals-to-balance-array/
package main

import "sort"

func minRemoval(nums []int, k int) (count int) {
	sort.Ints(nums)
	find := func(target int) int {
		l, r := 0, len(nums)
		for l < r {
			if mid := l + (r-l)/2; nums[mid] <= target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l - 1
	}
	for i, num := range nums {
		count = max(count, find(num*k)-i+1)
	}
	return len(nums) - count
}

func main() {}
