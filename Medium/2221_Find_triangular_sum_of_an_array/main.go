// https://leetcode.com/problems/find-triangular-sum-of-an-array/
package main

func triangularSum(nums []int) int {
	for ; len(nums) > 1; nums = nums[1:] {
		for i, prev := 1, nums[0]; i < len(nums); i++ {
			nums[i], prev = (nums[i]+prev)%10, nums[i]
		}
	}
	return nums[0]
}

func main() {}
