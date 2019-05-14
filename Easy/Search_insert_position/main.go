// https://leetcode.com/problems/search-insert-position/

package main

func searchInsert(nums []int, target int) int {
	length := len(nums)

	for i := 0; i < length; i++ {
		if target <= nums[i] {
			return i
		}
	}
	return length
}

func main() {}
