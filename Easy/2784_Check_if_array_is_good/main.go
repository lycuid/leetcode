// https://leetcode.com/problems/check-if-array-is-good/
package main

func isGood(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for nums[i] != i+1 {
			if nums[i] <= 0 || nums[i] >= len(nums) {
				return false
			}
			if nums[i] == nums[nums[i]-1] {
				if nums[i] == len(nums)-1 && nums[i] != nums[nums[i]] {
					nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
					continue
				}
				return false
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return nums[len(nums)-1] == len(nums)-1
}

func main() {}
