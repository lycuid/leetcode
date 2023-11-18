// https://leetcode.com/problems/frequency-of-the-most-frequent-element/
package main

import "sort"

func maxFrequency(nums []int, k int) int {
	used, max := 0, 1
	sort.Ints(nums)
	for i, j := 0, 1; j < len(nums); j++ {
		used += (nums[j] - nums[j-1]) * (j - i)
		for ; used > k && i < j; i++ {
			used -= nums[j] - nums[i]
		}
		if n := j - i + 1; n > max {
			max = n
		}
	}
	return max
}

func main() {}
