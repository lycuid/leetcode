// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
package main

import "sort"

func minPairSum(nums []int) (max int) {
	sort.Ints(nums)
	for i, n := 0, len(nums); i < n/2; i++ {
		if m := nums[i] + nums[n-i-1]; m > max {
			max = m
		}
	}
	return max
}

func main() {}
