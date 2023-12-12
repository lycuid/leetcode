// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
package main

import "sort"

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}

func main() {}
