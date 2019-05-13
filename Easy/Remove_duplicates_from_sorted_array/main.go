// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	length := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			length++
			nums[length-1] = nums[i]
		}
	}

	return length
}

func main() {}
