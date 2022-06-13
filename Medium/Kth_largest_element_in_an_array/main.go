// https://leetcode.com/problems/kth-largest-element-in-an-array/
package main

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func main() {}
