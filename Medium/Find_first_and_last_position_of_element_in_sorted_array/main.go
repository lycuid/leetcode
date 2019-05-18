// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

func getBounds(nums []int, pivot, start, end int) []int {
	lower, upper := pivot, pivot
	target := nums[pivot]

	if nums[start] == target {
		lower = start
	} else {
		for start < lower-1 {
			m := (start + lower) / 2
			if nums[m] == target {
				lower = m
			} else {
				start = m
			}
		}
	}

	if nums[end] == target {
		upper = end
	} else {
		for end > upper+1 {
			m := (end + upper) / 2
			if nums[m] == target {
				upper = m
			} else {
				end = m
			}
		}
	}

	return []int{lower, upper}
}

func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		n := nums[m]
		if n < target {
			l = m + 1
		} else if n > target {
			r = m - 1
		} else {
			return getBounds(nums, m, l, r)
		}
	}
	return []int{-1, -1}
}

func main() {}
