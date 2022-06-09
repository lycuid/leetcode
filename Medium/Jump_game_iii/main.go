// https://leetcode.com/problems/jump-game-iii/
package main

func Solve(start int, nums []int, visited []bool) bool {
	if start < 0 || start >= len(nums) || visited[start] {
		return false
	}
	if nums[start] == 0 {
		return true
	}
	visited[start] = true
	left, right := Solve(start-nums[start], nums, visited), Solve(start+nums[start], nums, visited)
	visited[start] = false
	return left || right
}

func canReach(nums []int, start int) bool {
	return Solve(start, nums, make([]bool, len(nums)))
}

func main() {}
