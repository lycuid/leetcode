// https://leetcode.com/problems/product-of-array-except-self/
package main

func productExceptSelf(nums []int) []int {
	cache, n := make([]int, len(nums)+1), len(nums)
	cache[n] = 1
	for i := n - 1; i >= 0; i-- {
		cache[i] = nums[i] * cache[i+1]
	}
	for i, prod := 0, 1; i < n; i++ {
		nums[i], prod = prod*cache[i+1], prod*nums[i]
	}
	return nums
}

func main() {}
