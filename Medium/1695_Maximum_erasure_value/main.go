// https://leetcode.com/problems/maximum-erasure-value/
package main

func maximumUniqueSubarray(nums []int) (res int) {
	var cache [1e4 + 1]bool
	for i, j, sum := 0, 0, 0; i < len(nums); i++ {
		for ; cache[nums[i]] && j <= i; j++ {
			sum, cache[nums[j]] = sum-nums[j], false
		}
		if sum, cache[nums[i]] = sum+nums[i], true; sum > res {
			res = sum
		}
	}
	return res
}

func main() {}
