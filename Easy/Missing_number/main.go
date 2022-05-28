// https://leetcode.com/problems/missing-number/
package main

func missingNumber(nums []int) (n int) {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i && nums[i] < len(nums) {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for ; n < len(nums) && n == nums[n]; n++ {
	}
	return n
}

func main() {}
