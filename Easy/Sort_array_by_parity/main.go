// https://leetcode.com/problems/sort-array-by-parity/
package main

func sortArrayByParity(nums []int) []int {
	cursor := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[cursor], nums[i] = nums[i], nums[cursor]
			cursor++
		}
	}
	return nums
}

func main() {}
