// https://leetcode.com/problems/product-of-array-except-self/
package main

func productExceptSelf(nums []int) []int {
	partial, result := make([]int, len(nums)+1), make([]int, len(nums))
	partial[len(nums)] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		partial[i] = partial[i+1] * nums[i]
	}
	for i, product := 0, 1; i < len(result); i++ {
		result[i] = product * partial[i+1]
		product *= nums[i]
	}
	return result
}

func main() {}
