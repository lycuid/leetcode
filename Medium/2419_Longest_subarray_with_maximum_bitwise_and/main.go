// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/

package main

func longestSubarray(nums []int) (size int) {
	if n := len(nums); n > 0 {
		maxnum := nums[0]
		for _, num := range nums {
			maxnum = max(maxnum, num)
		}
		for i, j := 0, 0; i < n; i = j + 1 {
			for j = i; j < n && nums[j] == maxnum; j++ {
			}
			size = max(size, j-i)
		}
	}
	return size
}

func main() {}
