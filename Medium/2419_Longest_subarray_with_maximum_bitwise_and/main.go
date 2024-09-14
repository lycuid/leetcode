// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
package main

func longestSubarray(nums []int) (count int) {
	maxnum := nums[0]
	for _, num := range nums {
		maxnum = max(maxnum, num)
	}
	for i, n := 0, len(nums); i < n; i++ {
		if nums[i] == maxnum {
			j := i
			for i+1 < n && nums[i+1] == maxnum {
				i++
			}
			count = max(count, i-j+1)
		}
	}
	return count
}

func main() {}
