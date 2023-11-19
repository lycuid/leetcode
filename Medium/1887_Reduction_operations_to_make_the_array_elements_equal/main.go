// https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
package main

import "sort"

func reductionOperations(nums []int) (count int) {
	sort.Ints(nums)
	for i, n := 1, len(nums); i < n; i++ {
		if nums[i] != nums[i-1] {
			count += n - i
		}
	}
	return count
}

func main() {}
