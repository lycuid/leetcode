// https://leetcode.com/problems/smallest-stable-index-i/
package main

import "math"

func firstStableIndex(nums []int, k int) int {
	cache := make([]int, len(nums)+1)
	cache[len(nums)] = math.MaxInt
	for i := len(nums) - 1; i >= 0; i-- {
		cache[i] = min(nums[i], cache[i+1])
	}
	var maxSoFar int
	for i, num := range nums {
		if maxSoFar = max(maxSoFar, num); maxSoFar-cache[i] <= k {
			return i
		}
	}
	return -1
}

func main() {}
