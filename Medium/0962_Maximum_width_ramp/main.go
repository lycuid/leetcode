// https://leetcode.com/problems/maximum-width-ramp/
package main

import "sort"

func maxWidthRamp(nums []int) (width int) {
	cache := make([]int, len(nums))
	for i := range cache {
		cache[i] = i
	}
	sort.Slice(cache, func(i, j int) bool {
		if nums[cache[i]] == nums[cache[j]] {
			return cache[i] < cache[j]
		}
		return nums[cache[i]] < nums[cache[j]]
	})
	min_index := cache[0]
	for i := range cache {
		min_index = min(min_index, cache[i])
		width = max(width, cache[i]-min_index)
	}
	return width
}

func main() {}
