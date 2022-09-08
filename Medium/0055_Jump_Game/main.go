// https://leetcode.com/problems/jump-game/
package main

func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := target - 1; i >= 0; i-- {
		if nums[i] >= target-i {
			target = i
		}
	}
	return target == 0
}

func main() {}
