// https://leetcode.com/problems/subarray-product-less-than-k/
package main

func numSubarrayProductLessThanK(nums []int, k int) (count int) {
	for i, j, prod := 0, 0, 1; j < len(nums); j++ {
		for prod *= nums[j]; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		count += j - i + 1
	}
	return count
}

func main() {}
