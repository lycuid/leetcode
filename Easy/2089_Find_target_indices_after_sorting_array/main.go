// https://leetcode.com/problems/find-target-indices-after-sorting-array/
package main

import "sort"

func targetIndices(nums []int, target int) (res []int) {
	sort.Ints(nums)
	left := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if left != len(nums) {
		right := sort.Search(len(nums), func(i int) bool {
			return nums[i] > target
		})
		for ; left < right; left++ {
			res = append(res, left)
		}
	}
	return res
}

func main() {}
