// https://leetcode.com/problems/maximum-product-of-three-numbers/
package main

import "sort"

func maximumProduct(nums []int) int {
	low, high := []int{nums[0], nums[1], nums[2]}, []int{nums[0], nums[1], nums[2]}
	sort.Ints(low)
	sort.Ints(high)
	for _, num := range nums[3:] {
		if num < low[0] {
			low[0], low[1], low[2] = num, low[0], low[1]
		} else if num < low[1] {
			low[1], low[2] = num, low[1]
		} else if num < low[2] {
			low[2] = num
		}
		if num > high[2] {
			high[2], high[1], high[0] = num, high[2], high[1]
		} else if num > high[1] {
			high[1], high[0] = num, high[1]
		} else if num > high[0] {
			high[0] = num
		}
	}
	return max(
		low[0]*low[1]*high[2],
		low[0]*low[1]*low[2],
		high[2]*high[1]*low[0],
		high[2]*high[1]*high[0],
	)
}

func main() {}
