// https://leetcode.com/problems/find-the-duplicate-number/
package main

func findDuplicate(nums []int) int {
	for i := range nums {
		for nums[i] != i+1 {
			if nums[nums[i]-1] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return 0
}

func main() {}
