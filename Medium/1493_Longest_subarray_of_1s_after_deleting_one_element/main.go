// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
package main

func longestSubarray(nums []int) (count int) {
	pre := make([]int, len(nums)+1)
	for i := range nums {
		if nums[i] == 1 {
			pre[i+1] = pre[i] + nums[i]
		}
	}
	nums = append(nums, 0)
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 1 {
			nums[i] += nums[i+1]
		}
		count = max(count, pre[i]+nums[i+1])
	}
	return count
}

func main() {}
