// https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/
package main

import "sort"

func missingInteger(nums []int) int {
	res, i := nums[0], 1
	for ; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		res += nums[i]
	}
	sort.Ints(nums[i:])
	for i -= 1; i < len(nums) && nums[i] <= res; i++ {
		if nums[i] == res {
			res++
		}
	}
	return res
}

func main() {}
