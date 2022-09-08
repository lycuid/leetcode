// https://leetcode.com/problems/rotate-array/
package main

func rotate(nums []int, k int) {
	l := len(nums)
	k %= l
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {}
