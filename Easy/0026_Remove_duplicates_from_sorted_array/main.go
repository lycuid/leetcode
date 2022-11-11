// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

func removeDuplicates(nums []int) (j int) {
	for i := 0; i < len(nums); j++ {
		for nums[j] = nums[i]; i < len(nums) && nums[i] == nums[j]; i++ {
		}
	}
	return j
}

func main() {}
