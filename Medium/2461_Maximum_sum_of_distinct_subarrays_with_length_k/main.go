// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
package main

func maximumSubarraySum(nums []int, k int) (ret int64) {
	cache, sum := make(map[int]int), int64(0)
	for i, j := 0, 0; j < len(nums); j++ {
		if index, found := cache[nums[j]]; found {
			for ; i <= index; i++ {
				sum -= int64(nums[i])
				delete(cache, nums[i])
			}
		}
		sum += int64(nums[j])
		cache[nums[j]] = j
		for ; j-i+1 > k; i++ {
			sum -= int64(nums[i])
			delete(cache, nums[i])
		}
		if j-i+1 == k {
			ret = max(ret, sum)
		}
	}
	return ret
}

func main() {}
