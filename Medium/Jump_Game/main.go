// https://leetcode.com/problems/jump-game/
package main

func maxint(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func canJump(nums []int) bool {
	l := len(nums)
	max := 0
	for i := 0; i < l; i++ {
		if i > max {
			return false
		}
		max = maxint(max, nums[i]+i)
	}

	return true
}

func main() {}
