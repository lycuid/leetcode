// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
package main

func maxProduct(nums []int) int {
	if nums[0] < nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	for _, num := range nums[2:] {
		if num > nums[0] {
			nums[0], nums[1] = num, nums[0]
		} else if num > nums[1] {
			nums[1] = num
		}
	}
	return (nums[0] - 1) * (nums[1] - 1)
}

func main() {}
