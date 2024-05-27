// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
package main

import "sort"

func specialArray(nums []int) int {
	if n := len(nums); n > 0 {
		if sort.Ints(nums); nums[0] >= n {
			return n
		}
		for i := 1; i <= len(nums); i++ {
			if nums[n-i] >= i && nums[n-i-1] < i {
				return i
			}
		}
	}
	return -1
}

func main() {}
