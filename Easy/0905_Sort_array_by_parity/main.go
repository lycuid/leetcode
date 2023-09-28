// https://leetcode.com/problems/sort-array-by-parity/
package main

func sortArrayByParity(nums []int) []int {
	for l, r := 0, len(nums)-1; l < r; l++ {
		for ; l < r && nums[l]%2 != 0; r-- {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return nums
}

func main() {}
