// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/
package main

func maxSubarrayLength(nums []int, k int) (count int) {
	cache := make(map[int]int)
	for i, j := 0, 0; j < len(nums); j++ {
		for cache[nums[j]]++; cache[nums[j]] > k; i++ {
			cache[nums[i]]--
		}
		if n := j - i + 1; n > count {
			count = n
		}
	}
	return count
}

func main() {}
