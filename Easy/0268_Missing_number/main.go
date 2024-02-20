// https://leetcode.com/problems/missing-number/
package main

func missingNumber(nums []int) int {
	n := len(nums)
	for i := range nums {
		for nums[i] != i && nums[i] != n {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return n
}

func main() {}
