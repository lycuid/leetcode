// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
package main

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	for i := 0; i <= len(nums)-k; i++ {
		nums[0] = min(nums[0], nums[i]-nums[i+k-1])
	}
	return nums[0]
}

func main() {}
