// https://leetcode.com/problems/jump-game/
package main

func canJump(nums []int) bool {
	max, n := 0, len(nums)-1
	for i := 0; i <= n && i <= max; i++ {
		if m := i + nums[i]; m > max {
			max = m
		}
	}
	return max >= n
}

func main() {}
