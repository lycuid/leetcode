// https://leetcode.com/problems/longest-subsequence-with-limited-sum/
package main

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for i := 0; i < len(queries); i++ {
		l, r := 0, len(nums)-1
		for l < r {
			if mid := (l + r) / 2; nums[mid] < queries[i] {
				l = mid + 1
			} else if nums[mid] > queries[i] {
				r = mid
			} else {
				l, r = mid, mid
			}
		}
		for ; l < len(nums) && nums[l] <= queries[i]; l++ {
		}
		queries[i] = l
	}
	return queries
}

func main() {}
