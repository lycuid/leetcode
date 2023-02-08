// https://leetcode.com/problems/jump-game-ii/
package main

func jump(nums []int) int {
	cache, n := make([]int, len(nums)), len(nums)-1
	for i, j := 0, 1; i <= n && j <= n; i++ {
		for ; j <= i+nums[i] && j <= n; j++ {
			cache[j] = cache[i] + 1
		}
	}
	return cache[n]
}

func main() {}
