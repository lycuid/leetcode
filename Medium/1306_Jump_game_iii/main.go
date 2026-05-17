// https://leetcode.com/problems/jump-game-iii/
package main

func find(start int, nums []int, visited []bool) bool {
	if start < 0 || start >= len(nums) || visited[start] {
		return false
	}
	if nums[start] == 0 {
		return true
	}
	visited[start] = true
	l := find(start-nums[start], nums, visited)
	r := find(start+nums[start], nums, visited)
	visited[start] = false
	return l || r
}

func canReach(arr []int, start int) bool {
	return find(start, arr, make([]bool, len(arr)))
}

func main() {}
