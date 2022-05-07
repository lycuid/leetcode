// https://leetcode.com/problems/132-pattern/
package main

import "math"

func find132pattern(nums []int) bool {
	cursor, nums_k := 0, math.MinInt
	stack := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < nums_k {
			return true
		}
		for cursor > 0 && stack[cursor-1] < nums[i] {
			nums_k = stack[cursor-1]
			cursor--
		}
		stack[cursor] = nums[i]
		cursor++
	}
	return false
}

func main() {}
