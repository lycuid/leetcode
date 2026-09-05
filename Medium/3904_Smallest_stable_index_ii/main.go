// https://leetcode.com/problems/smallest-stable-index-ii/
package main

import "math"

func firstStableIndex(nums []int, k int) int {
	cache, n := make([]int, len(nums)+1), len(nums)
	cache[n] = math.MaxInt
	for i := n - 1; i >= 0; i-- {
		cache[i] = min(nums[i], cache[i+1])
	}
	cache[n] = 0
	for i, num := range nums {
		if cache[n] = max(cache[n], num); cache[n]-cache[i] <= k {
			return i
		}
	}
	return -1
}

func main() {}
