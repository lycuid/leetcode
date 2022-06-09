// https://leetcode.com/problems/jump-game-ii/
package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func jump(nums []int) int {
	n := len(nums)
	min := make([]int, n)
	for i, j := 0, 1; i < j && j < n; i++ {
		for ; j <= Min(n-1, i+nums[i]); j++ {
			min[j] = min[i] + 1
		}
	}
	return min[n-1]
}

func main() {}
