// https://leetcode.com/problems/house-robber/
package main

func rob(nums []int) int {
	for i := range nums {
		if i > 1 && nums[i] < nums[i-2]+nums[i] {
			nums[i] = nums[i-2] + nums[i]
		}
		if i > 0 && nums[i] < nums[i-1] {
			nums[i] = nums[i-1]
		}
	}
	return nums[len(nums)-1]
}

func main() {}
