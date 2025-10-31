// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/
package main

func getSneakyNumbers(nums []int) []int {
	for i, n := 0, len(nums)-1; i < len(nums)-2; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				nums[i], nums[n], n = nums[n], nums[i], n-1
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return nums[len(nums)-2:]
}

func main() {}
