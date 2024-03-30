// https://leetcode.com/problems/subarrays-with-k-different-integers/
package main

func Aux(nums []int, k int) (count int) {
	cache := make(map[int]int)
	for i, j := 0, 0; j < len(nums); j++ {
		for cache[nums[j]]++; i <= j && len(cache) > k; i++ {
			if cache[nums[i]]--; cache[nums[i]] == 0 {
				delete(cache, nums[i])
			}
		}
		count += j - i + 1
	}
	return count
}

func subarraysWithKDistinct(nums []int, k int) int {
	return Aux(nums, k) - Aux(nums, k-1)
}

func main() {}
