// https://leetcode.com/problems/matchsticks-to-square/
package main

import "sort"

func dfs(nums []int, i, s1, s2, s3, s4 int) bool {
	if i == len(nums) {
		return s1 == 0 && s2 == 0 && s3 == 0 && s4 == 0
	}
	if nums[i] <= s1 && dfs(nums, i+1, s1-nums[i], s2, s3, s4) {
		return true
	}
	if nums[i] <= s2 && dfs(nums, i+1, s1, s2-nums[i], s3, s4) {
		return true
	}
	if nums[i] <= s3 && dfs(nums, i+1, s1, s2, s3-nums[i], s4) {
		return true
	}
	if nums[i] <= s4 && dfs(nums, i+1, s1, s2, s3, s4-nums[i]) {
		return true
	}
	return false
}

func makesquare(nums []int) bool {
	var sum int
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	for _, num := range nums {
		sum += num
	}
	if sum%4 != 0 {
		return false
	}
	return dfs(nums, 0, sum/4, sum/4, sum/4, sum/4)
}

func main() {}
