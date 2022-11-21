// https://leetcode.com/problems/132-pattern/
package main

func find132pattern(nums []int) bool {
	stack, cursor, last := make([]int, len(nums)), 0, int(-(1e9 + 7))
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < last {
			return true
		}
		for ; cursor > 0 && stack[cursor-1] < nums[i]; cursor-- {
			last = stack[cursor-1]
		}
		stack[cursor], cursor = nums[i], cursor+1
	}
	return false
}

func main() {}
