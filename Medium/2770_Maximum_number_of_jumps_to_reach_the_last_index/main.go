// https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index/
package main

func maximumJumps(nums []int, target int) int {
	cache := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		cache[i] = -1
		for j := i - 1; j >= 0; j-- {
			if diff := nums[i] - nums[j]; cache[j] != -1 && diff >= -target && diff <= target {
				cache[i] = max(cache[i], cache[j]+1)
			}
		}
	}
	return cache[len(nums)-1]
}

func main() {}
