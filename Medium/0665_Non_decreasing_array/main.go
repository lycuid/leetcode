// https://leetcode.com/problems/non-decreasing-array/
package main

import "math"

func checkPossibility(nums []int) (modified bool) {
	nums = append([]int{math.MinInt}, append(nums, math.MaxInt)...)
	for i := 1; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			if modified || (nums[i-1] > nums[i+1] && nums[i] > nums[i+2]) {
				return false
			}
			modified = true
		}
	}
	return true
}

func main() {}
