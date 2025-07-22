// https://leetcode.com/problems/maximum-erasure-value/
package main

func maximumUniqueSubarray(nums []int) (count int) {
	cache, current := make(map[int]bool), 0
	for i, j := 0, 0; i < len(nums); i++ {
		for ; cache[nums[i]] && j <= i; j++ {
			cache[nums[j]], current = false, current-nums[j]
		}
		cache[nums[i]], current = true, current+nums[i]
		count = max(count, current)
	}
	return count
}

func main() {}
