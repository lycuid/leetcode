// https://leetcode.com/problems/product-of-array-except-self/
package main

func productExceptSelf(nums []int) []int {
	l := len(nums)

	products := make([]int, l)
	products[l-1] = nums[l-1]
	for i := l - 2; i > 0; i-- {
		products[i] = nums[i] * products[i+1]
	}

	acc := 1
	for i := 0; i < l-1; i++ {
		products[i] = acc * products[i+1]
		acc *= nums[i]
	}
	products[l-1] = acc

	return products
}

func main() {}
