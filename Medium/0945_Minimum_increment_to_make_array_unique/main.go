// https://leetcode.com/problems/minimum-increment-to-make-array-unique/
package main

import "sort"

func minIncrementForUnique(nums []int) (count int) {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			count += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return count
}

func main() {}
