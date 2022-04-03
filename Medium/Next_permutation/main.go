// https://leetcode.com/problems/next-permutation/
package main

import "sort"

func find_smallest_in_range(nums []int, min, max, start, end int) int {
	ret := start
	for i := start; i <= end; i++ {
		if nums[i] > min && nums[i] <= max {
			ret = i
		}
	}
	return ret
}

func nextPermutation(nums []int) {
	l, i := len(nums), 0
	if l <= 0 {
		return
	}
	i, h := l-2, l-1
	for ; i >= 0; i-- {
		if nums[i] > nums[h] {
			h = i
		}
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		to_swap := find_smallest_in_range(nums, nums[i], nums[h], i, l-1)
		nums[to_swap], nums[i] = nums[i], nums[to_swap]
	}
	sort.Ints(nums[i+1:])
}

func main() {}
