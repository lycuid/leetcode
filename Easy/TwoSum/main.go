// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	bitmap := make(map[int]int)
	for i := range nums {
		bitmap[nums[i]] = i
	}

	for v1 := range nums {
		c := target - nums[v1]
		if v2, ok := bitmap[c]; v1 != v2 && ok {
			return []int{v1, v2}
		}
	}
	return []int{}
}

