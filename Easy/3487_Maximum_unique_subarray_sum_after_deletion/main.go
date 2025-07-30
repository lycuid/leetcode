// https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/
package main

import "sort"

func maxSum(nums []int) (sum int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum = nums[0]
	for i, prev := 1, sum; i < len(nums); i, prev = i+1, nums[i] {
		newSum := sum + nums[i]
		if newSum < sum {
			break
		}
		if prev != nums[i] {
			sum = newSum
		}
	}
	return sum
}

func main() {}
