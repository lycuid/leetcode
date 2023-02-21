// https://leetcode.com/problems/single-element-in-a-sorted-array/
package main

func singleNonDuplicate(nums []int) (l int) {
	for r := len(nums) - 1; l < r; {
		if m := (l + r) / 2; nums[m] == nums[m+1-(m%2)*2] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

func main() {}
