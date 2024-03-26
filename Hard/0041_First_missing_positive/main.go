// https://leetcode.com/problems/first-missing-positive/
package main

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] < len(nums) && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	ret := 1
	for ret <= len(nums) && nums[ret-1] == ret {
		ret++
	}
	return ret
}

func main() {}
