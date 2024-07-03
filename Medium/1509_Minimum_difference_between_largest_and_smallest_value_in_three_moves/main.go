// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
package main

import "sort"

func minDifference(nums []int) (diff int) {
	if n := len(nums); n > 3 {
		sort.Ints(nums)
		diff = min(nums[n-1]-nums[3], nums[n-2]-nums[2], nums[n-3]-nums[1], nums[n-4]-nums[0])
	}
	return diff
}

func main() {}
