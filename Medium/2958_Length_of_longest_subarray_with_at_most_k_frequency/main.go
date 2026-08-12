// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/
package main

func maxSubarrayLength(nums []int, k int) (res int) {
	cache := make(map[int]int)
	for i, j := 0, 0; i < len(nums); i++ {
		for cache[nums[i]]++; cache[nums[i]] > k; j++ {
			cache[nums[j]]--
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {}
