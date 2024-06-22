// https://leetcode.com/problems/count-number-of-nice-subarrays/
package main

func numberOfSubarrays(nums []int, k int) (ret int) {
	cache := make([]int, len(nums)+1)
	cache[0] = 1
	for i, t := 0, 0; i < len(nums); i++ {
		if t += nums[i] & 1; t-k >= 0 {
			ret += cache[t-k]
		}
		cache[t]++
	}
	return ret
}

func main() {}
