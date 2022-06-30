// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/
package main

import "sort"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMoves2(nums []int) (ret int) {
	sort.Ints(nums)
	n := nums[len(nums)/2]
	for i := range nums {
		ret += Abs(nums[i] - n)
	}
	return ret
}

func main() {}
